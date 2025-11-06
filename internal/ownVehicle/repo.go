package ownvehicle

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TripLite struct {
	ID          uint `gorm:"column:id"`
    LoadPoint   *string  `gorm:"column:load_point"`
    UnloadPoint *string  `gorm:"column:unload_point"`
    Advance     *float64 `gorm:"column:advance"`
}

// Repository abstracts DB access for OwnVehicle use cases.
type Repository interface {
    // Fetch trips for a driver between UTC times (inclusive).
    GetTripsByDriverBetween(driverName string, startUTC, endUTC time.Time) ([]TripLite, error)
    // Lookup rate for a (load_point, unload_point) pair.
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
    // Use created_at (timestamptz stored in UTC) for accurate window filtering
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
    // INSERT ... ON CONFLICT (trip_id) DO NOTHING
    return r.db.Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "trip_id"}},
        DoNothing: true,
    }).Create(item).Error
}