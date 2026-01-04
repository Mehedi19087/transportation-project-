package dailysidecash

import (
	"time"
	"errors"

	"gorm.io/gorm"
)

type DailySideCashRepo interface {
	Create(dailySideCash *DailySideCash) error
	GetByID(id uint) (*DailySideCash, error)
	GetAll(productID uint, limit, offset int) ([]DailySideCash, int64, error)
	Update(dailySideCash *DailySideCash) error
	Delete(id uint) error
	GetByDate(productID uint, date time.Time) (*DailySideCash, error)
}

type dailySideCashRepo struct {
	 db *gorm.DB 
}

func NewDailySideCashRepo(db *gorm.DB) DailySideCashRepo {
	 return &dailySideCashRepo {db : db}
}

func (r *dailySideCashRepo) Create(dailySideCash *DailySideCash) error {
	 return r.db.Create(dailySideCash).Error
}

func(r *dailySideCashRepo) GetByID(id uint) (*DailySideCash, error) {
	 var res DailySideCash 

	 if err := r.db.First(&res, id).Error; err!= nil {
		 return nil, err 
	 }

	 return &res, nil 
}


func(r *dailySideCashRepo) GetAll(productID uint, offset, limit int) ([]DailySideCash,int64, error) {

	var dailySideCash []DailySideCash 
	var total int64 

	query := r.db.Model(&DailySideCash{}).Where("product_id = ?", productID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err 
	}
	 if err:= query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&dailySideCash).Error; err!= nil {
		 return nil , 0 , err 
	 }
	 return dailySideCash, total, nil 
}

func (r *dailySideCashRepo) Update(dailySideCash *DailySideCash) error {
   return r.db.Save(dailySideCash).Error
}

func(r *dailySideCashRepo) Delete(id uint) error {
	 return r.db.Delete(&DailySideCash{},id).Error
}


func (r *dailySideCashRepo) GetByDate(productID uint, date time.Time) (*DailySideCash, error) {
   var dailySideCash DailySideCash

   // Use strictly the date part for comparison to avoid time issues if needed, 
   // but assuming the input 'date' is already truncated or the DB type handles it.
   // Using formatted string to be safe with DATE types in Postgres.
   dateStr := date.Format("2006-01-02")

   err := r.db.Where("product_id = ? AND date = ?", productID, dateStr).First(&dailySideCash).Error 

   if err != nil {
          if errors.Is(err, gorm.ErrRecordNotFound) {
              return nil, errors.New("daily side cash record not found for the given date and product")
          }
          return nil, err
      }
      return &dailySideCash, nil

}