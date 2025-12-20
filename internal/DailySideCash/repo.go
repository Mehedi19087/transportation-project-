package dailysidecash

import (
	"time"
	"errors"

	"gorm.io/gorm"
)

type DailySideCashRepo interface {
	Create(dailySideCash *DailySideCash) error
	GetByID(id uint) (*DailySideCash, error)
	GetAll(limit, offset int) ([]DailySideCash, int64, error)
	Update(dailySideCash *DailySideCash) error
	Delete(id uint) error
	GetByDate(date time.Time) (*DailySideCash, error)
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


func(r *dailySideCashRepo) GetAll(offset, limit int) ([]DailySideCash,int64, error) {

	var dailySideCash []DailySideCash 
	var total int64 

	if err := r.db.Model(&DailySideCash{}).Count(&total).Error; err != nil {
		return nil, 0, err 
	}
	 if err:= r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&dailySideCash).Error; err!= nil {
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


func (r *dailySideCashRepo) GetByDate(date time.Time) (*DailySideCash, error) {
   var dailySideCash DailySideCash

   err := r.db.Where("date = ?", date.Format("2006-01,02")).First(&dailySideCash).Error 

   if err != nil {
          if errors.Is(err, gorm.ErrRecordNotFound) {
              return nil, errors.New("daily side cash record not found for the given date")
          }
          return nil, err
      }
      return &dailySideCash, nil

}