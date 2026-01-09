package auth

import (
	"errors"
	"transportation/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	 CreateUser(req *AuthReq) error
	 Login(req *AuthReq) (string, string, error)
	 UpdateUser(id uint, req UpdateReq, requesterRole string) error
	 GetPendingUsers() ([]User, error)
}

type authService struct {
	 repo AuthRepo
}

func NewAuthService (repo AuthRepo) AuthService {
	 return &authService {repo: repo}
}

func(s *authService) CreateUser(req *AuthReq) error {
	 if req.Name == "" {
		 return errors.New("name is required")
	 }
	 if req.Password == "" {
		 return errors.New("password is required")
	 }
	 if _, err := s.repo.GetUserByName(req.Name) ; err == nil {
		 return errors.New("username already exists")
	 }

	 hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	 if err != nil {
           return err
     }

	 user := &User {
		Name : req.Name,
		Password: string(hashedPassword),
	 }

	 if req.ProductID != nil {
		 user.ProductID = *req.ProductID
	 }

	 if err := s.repo.CreateUser(user); err != nil {
		 return err 
	 }
	 return nil 
}

func (s *authService) Login(req *AuthReq) (string, string, error) {
	  if req.Name == "" || req.Password == "" {
           return "", "", errors.New("username and password are required")
      }
	  user, err := s.repo.GetUserByName(req.Name)
       if err != nil {
           return "", "", errors.New("invalid username or password")
      }
	  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
           return "", "", errors.New("invalid username or password")
      }
	  if user.Status != "active" {
           return "", "", errors.New("account is pending approval")
      }
	  token, err := utils.GenerateJWT(user.ID, user.Name, user.Role, &user.ProductID)
	  return token, user.Role, err
}

func (s *authService) UpdateUser(id uint, req UpdateReq, requesterRole string) error {
       // 1. We don't strictly need to fetch the user first if we trust the ID, 
       //    but checking existence is good practice.
       _, err := s.repo.GetUserByID(id)
       if err != nil {
           return errors.New("user not found")
       }

       // 2. Create a map to hold ONLY the fields we want to change
       updates := make(map[string]interface{})

       if req.Name != nil && *req.Name != "" {
           updates["name"] = *req.Name
       }

       if req.Password != nil && *req.Password != "" {
           hashed, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
           if err != nil {
               return err
           }
           updates["password"] = string(hashed)
       }

       if req.Status != nil {
           if requesterRole != RoleAdmin {
               return errors.New("unauthorized")
           }
           updates["status"] = *req.Status
       }
       return s.repo.UpdateUserRecord(id, updates)
   }
func (s *authService) GetPendingUsers() ([]User, error) {
       return s.repo.GetPendingUsers()
}