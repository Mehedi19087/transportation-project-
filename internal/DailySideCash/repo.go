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
	GetLastRecordBeforeDate(productID uint, date time.Time) (*DailySideCash, error)
	GetFirstRecordAfterDate(productID uint, date time.Time) (*DailySideCash, error)
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

	query := r.db.Model(&DailySideCash{})
	if productID != 0 {
		query = query.Where("product_id = ?", productID)
	}

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

   // Format the date as a string YYYY-MM-DD to match against the date column
   // We cast the DB column to DATE to ignore the time component if it exists
   dateStr := date.Format("2006-01-02")

   err := r.db.Where("product_id = ? AND DATE(date) = ?", productID, dateStr).First(&dailySideCash).Error 

   if err != nil {
          if errors.Is(err, gorm.ErrRecordNotFound) {
              return nil, errors.New("daily side cash record not found for the given date and product")
          }
          return nil, err
      }
      return &dailySideCash, nil

}

func (r *dailySideCashRepo) GetLastRecordBeforeDate(productID uint, date time.Time) (*DailySideCash, error) {
	var dailySideCash DailySideCash
	dateStr := date.Format("2006-01-02")
	err := r.db.Where("product_id = ? AND DATE(date) < ?", productID, dateStr).Order("date desc").First(&dailySideCash).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no previous record found
		}
		return nil, err
	}
	return &dailySideCash, nil
}

func (r *dailySideCashRepo) GetFirstRecordAfterDate(productID uint, date time.Time) (*DailySideCash, error) {
	var dailySideCash DailySideCash
	dateStr := date.Format("2006-01-02")
	err := r.db.Where("product_id = ? AND DATE(date) > ?", productID, dateStr).Order("date asc").First(&dailySideCash).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no next record found
		}
		return nil, err
	}
	return &dailySideCash, nil
}