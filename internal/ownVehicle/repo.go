package ownvehicle

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TripLite struct {
	ID          uint     `gorm:"column:id"`
	LoadPoint   *string  `gorm:"column:load_point"`
	UnloadPoint *string  `gorm:"column:unload_point"`
	Advance     *float64 `gorm:"column:advance"`
}

// Repository abstracts DB access for legacy OwnVehicle use cases.
type Repository interface {
	GetTripsByDriverBetween(driverName string, startUTC, endUTC time.Time) ([]TripLite, error)
	GetRateByLocations(loadPoint, unloadPoint string) (float64, error)
	CreateOwnVehicle(item *OwnVehicle) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) GetTripsByDriverBetween(driverName string, startUTC, endUTC time.Time) ([]TripLite, error) {
	var trips []TripLite
	err := r.db.Table("trips").
		Select("id, load_point, unload_point, advance").
		Where("driver_name = ? AND created_at >= ? AND created_at <= ?", driverName, startUTC, endUTC).
		Order("created_at ASC").
		Find(&trips).Error
	return trips, err
}

func (r *repo) GetRateByLocations(loadPoint, unloadPoint string) (float64, error) {
	var row struct {
		Rate float64 `gorm:"column:rate"`
	}
	err := r.db.Table("route_pricings").
		Select("rate").
		Where("load_point = ? AND unload_point = ?", loadPoint, unloadPoint).
		First(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("rate not found for given locations")
		}
		return 0, err
	}
	return row.Rate, nil
}

func (r *repo) CreateOwnVehicle(item *OwnVehicle) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "trip_id"}},
		DoNothing: true,
	}).Create(item).Error
}

type OwnVehicleTripRepository interface {
	Create(ownVehicleTrip *OwnVehicleTrip) error
	Get(id uint) (*OwnVehicleTrip, error)
	GetAll(offset, limit int) ([]OwnVehicleTrip, int64, error)
	Update(ownVehicleTrip *OwnVehicleTrip) error
	Delete(id uint) error
	GetByVehicleAndDateRange(vehicleNo string, startDate, endDate string) ([]OwnVehicleTrip, error)
    ExistsByTripID(tripID uint) (bool, error)
    GetByTripID(tripID uint) (*OwnVehicleTrip, error)
    CreateWithTx(tx *gorm.DB, ownVehicleTrip *OwnVehicleTrip) error
    UpdateWithTx(tx *gorm.DB, ownVehicleTrip *OwnVehicleTrip) error
    DeleteByTripIDWithTx(tx *gorm.DB, tripID uint) error
    GetDB() *gorm.DB
}

type ownVehicleTripRepository struct {
	db *gorm.DB
}

func NewOwnVehicleTripRepository(db *gorm.DB) OwnVehicleTripRepository {
	return &ownVehicleTripRepository{db: db}
}

func (r *ownVehicleTripRepository) Create(ownVehicleTrip *OwnVehicleTrip) error {
	return r.db.Create(ownVehicleTrip).Error
}

func (r *ownVehicleTripRepository) Get(id uint) (*OwnVehicleTrip, error) {
	var trip OwnVehicleTrip
	if err := r.db.First(&trip, id).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *ownVehicleTripRepository) GetAll(offset, limit int) ([]OwnVehicleTrip, int64, error) {
	var total int64
	var trips []OwnVehicleTrip

	if err := r.db.Model(&OwnVehicleTrip{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Order("date DESC, created_at DESC").
		Limit(limit).Offset(offset).
		Find(&trips).Error; err != nil {
		return nil, 0, err
	}

	return trips, total, nil
}

func (r *ownVehicleTripRepository) Update(ownVehicleTrip *OwnVehicleTrip) error {
	return r.db.Save(ownVehicleTrip).Error
}

func (r *ownVehicleTripRepository) Delete(id uint) error {
	return r.db.Delete(&OwnVehicleTrip{}, id).Error
}

func (r *ownVehicleTripRepository) GetByVehicleAndDateRange(vehicleNo string, startDate, endDate string) ([]OwnVehicleTrip, error) {
	var trips []OwnVehicleTrip
	if err := r.db.Where("vehicle_no = ? AND date >= ? AND date <= ?", vehicleNo, startDate, endDate).
		Order("date ASC, created_at ASC").
		Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *ownVehicleTripRepository) ExistsByTripID(tripID uint) (bool, error) {
    var count int64
    err := r.db.Model(&OwnVehicleTrip{}).Where("trip_id = ?", tripID).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *ownVehicleTripRepository) GetByTripID(tripID uint) (*OwnVehicleTrip, error) {
    var trip OwnVehicleTrip
    if err := r.db.Where("trip_id = ?", tripID).First(&trip).Error; err != nil {
        return nil, err
    }
    return &trip, nil
}

func (r *ownVehicleTripRepository) CreateWithTx(tx *gorm.DB, ownVehicleTrip *OwnVehicleTrip) error {
    return tx.Create(ownVehicleTrip).Error
}

func (r *ownVehicleTripRepository) UpdateWithTx(tx *gorm.DB, ownVehicleTrip *OwnVehicleTrip) error {
    return tx.Save(ownVehicleTrip).Error
}

func (r *ownVehicleTripRepository) GetDB() *gorm.DB {
    return r.db
}

func (r *ownVehicleTripRepository) DeleteByTripIDWithTx(tx *gorm.DB, tripID uint) error {
    return tx.Where("trip_id = ?", tripID).Delete(&OwnVehicleTrip{}).Error
}