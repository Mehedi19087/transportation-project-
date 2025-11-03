package dealer 

import (
	"gorm.io/gorm"
	"errors"
)


type DealerRepo interface {
	 Create(dealer *Dealer) error 
	 Update(dealer *Dealer) error 
	 GetAll(offset , limit int) ([]Dealer,int64,error)
	 Get(id uint) (*Dealer, error)
	 Delete(id uint) error
}

type dealerRepo struct {
	 db *gorm.DB 
}

func NewDealerRepo(db *gorm.DB) DealerRepo{
	return &dealerRepo{db:db}
}

func(r *dealerRepo) Create(dealer *Dealer) error {
	return r.db.Create(dealer).Error
}

func(r *dealerRepo) Update(dealer *Dealer) error {
	 return r.db.Save(dealer).Error 
}

func(r *dealerRepo) Get(id uint) (*Dealer, error) {
	var res Dealer 
	err := r.db.First(&res, id).Error 
	if err != nil {
		if err == gorm.ErrRecordNotFound {
            return nil, errors.New("dealer not found")
        }
        return nil, err 
	}
	return &res , nil 
}

func (r *dealerRepo) Delete(id uint) error {
	 return r.db.Delete(&Dealer{}, id).Error 
} 

func(r *dealerRepo) GetAll(offset, limit int) ([]Dealer,int64, error) {
     var dealers []Dealer 
	 var total int64 
	 if err := r.db.Model(&Dealer{}).Count(&total).Error; err != nil {
		 return nil, 0, err 
	 }
	 if err:= r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&dealers).Error; err!= nil {
		 return nil , 0 , err 
	 }
	 return dealers, total, nil 
}