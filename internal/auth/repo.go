package auth

import "gorm.io/gorm"

type AuthRepo interface {
	CreateUser(user *User) error 
	GetUserByName(name string) (*User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUserRecord(id uint, updates map[string]interface{}) error
	GetPendingUsers() ([]User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo (db *gorm.DB) AuthRepo {
	 return &authRepo {db : db}
}

func(r *authRepo) CreateUser(user *User) error {
	 return r.db.Create(user).Error 
}

func(r *authRepo) GetUserByName(name string) (*User, error) {
	 var user User 
	 if err := r.db.Where("name = ?", name).First(&user).Error ; err != nil {
		 return nil, err 
	 }
	 return &user , nil 
}

func (r *authRepo) GetUserByID(id uint) (*User, error) {
       var user User
       if err := r.db.First(&user, id).Error; err != nil {
           return nil, err
       }
       return &user, nil
   }

func (r *authRepo) UpdateUserRecord(id uint, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *authRepo) GetPendingUsers() ([]User, error) {
       var users []User
       if err := r.db.Where("status = ?", "pending").Find(&users).Error; err != nil {
           return nil, err
       }
       return users, nil
}
