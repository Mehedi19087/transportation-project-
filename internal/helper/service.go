package helper

import (
    "errors"
    "fmt"

    "gorm.io/gorm"
)

type HelperService interface {
    CreateHelper(helper *CreateHelperDTO) (*Helper, error)
    GetHelper(id uint) (*Helper, error)
    UpdateHelper(id uint, req *UpdateHelperDTO) error
    DeleteHelper(id uint) error
    GetAllHelper(page, pageSize int) ([]Helper, int64, error)
}

type helperService struct {
    helperRepo HelperRepo
}

func NewHelperService(helperRepo HelperRepo) HelperService {
    return &helperService{helperRepo: helperRepo}
}

func (s *helperService) CreateHelper(req *CreateHelperDTO) (*Helper, error) {
    if req.Mobile == "" || req.Name == "" {
        return nil, errors.New("name and mobile are required")
    }

    exists, err := s.helperRepo.ExistsByMobile(req.Mobile)
    if err != nil {
        return nil, err 
    }
    if exists {
        return nil, errors.New("helper with this mobile already exists")
    }

    helper := &Helper{
        Name:            req.Name,
        Mobile:          req.Mobile,
        Emergency:       req.Emergency,
        Address:         req.Address,
        Salary:          req.Salary,
        AssignedVehicle: req.AssignedVehicle,
        JoiningDate:     req.JoiningDate,
        Experience:      req.Experience,
        NID:             req.NID,
        Status:          req.Status,
        Image:           req.Image,
    }

    if helper.Status == "" {
        helper.Status = "Active"
    }

    if err := s.helperRepo.Create(helper); err != nil {
        return nil, err 
    }

    return helper, nil
}

func (s *helperService) GetHelper(id uint) (*Helper, error) {
    helper, err := s.helperRepo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("helper not found")
        }
        return nil, err
    }
    return helper, nil
}

func (s *helperService) UpdateHelper(id uint, req *UpdateHelperDTO) error {
    helper, err := s.helperRepo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("helper not found")
        }
        return err
    }

    // Update only provided fields
    if req.Name != nil {
        helper.Name = *req.Name
    }
    if req.Mobile != nil {
        helper.Mobile = *req.Mobile
    }
    if req.Emergency != nil {
        helper.Emergency = *req.Emergency
    }
    if req.Address != nil {
        helper.Address = *req.Address
    }
    if req.Salary != nil {
        helper.Salary = *req.Salary
    }
    if req.AssignedVehicle != nil {
        helper.AssignedVehicle = *req.AssignedVehicle
    }
    if req.JoiningDate != nil {
        helper.JoiningDate = req.JoiningDate
    }
    if req.Experience != nil {
        helper.Experience = *req.Experience
    }
    if req.NID != nil {
        helper.NID = *req.NID
    }
    if req.Status != nil {
        helper.Status = *req.Status
    }
    if req.Image != nil {
        helper.Image = *req.Image
    }

    if err := s.helperRepo.Update(helper); err != nil {
        return fmt.Errorf("failed to update helper: %w", err)
    }

    return nil
}

func (s *helperService) DeleteHelper(id uint) error {
    _, err := s.helperRepo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("helper not found")
        }
        return err
    }

    if err := s.helperRepo.Delete(id); err != nil {
        return fmt.Errorf("failed to delete helper: %w", err)
    }

    return nil
}

func (s *helperService) GetAllHelper(page, pageSize int) ([]Helper, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    helpers, total, err := s.helperRepo.GetAll(offset, pageSize)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to get helpers: %w", err)
    }

    return helpers, total, nil
}