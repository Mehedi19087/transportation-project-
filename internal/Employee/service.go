package employee

import (
    "errors"
    "fmt"
    "time"

    "gorm.io/gorm"
)

type EmployeeService interface {
    CreateEmployee(req *CreateEmployeeReq) error
    GetEmployee(id uint) (*Employee, error)
    UpdateEmployee(id uint, req *UpdateEmployeeReq) error
    DeleteEmployee(id uint) error
    GetAllEmployee(page, pageSize int) ([]Employee, int64, error)
}

type employeeService struct {
    repo EmployeeRepo
}

func NewEmployeeService(repo EmployeeRepo) EmployeeService {
    return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(req *CreateEmployeeReq) error {
    employee := &Employee{
        FullName:       &req.FullName,
        Email:          req.Email,
        Mobile:         req.Mobile,
        BirthDate:      req.BirthDate,
        Gender:         req.Gender,
        BloodGroup:     req.BloodGroup,
        Designation:    req.Designation,
        EmploymentType: req.EmploymentType,
        BranchName:     req.BranchName,
        JoinDate:       req.JoinDate,
        Salary:         req.Salary,
        PaymentMethod:  req.PaymentMethod,
        Nid:            req.Nid,
        Address:        req.Address,
        Status:         req.Status,
        Image:          req.Image,
        Role:           req.Role,
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
    }

    if err := s.repo.Create(employee); err != nil {
        return fmt.Errorf("failed to create employee: %w", err)
    }

    return nil
}

func (s *employeeService) GetEmployee(id uint) (*Employee, error) {
    res, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *employeeService) UpdateEmployee(id uint, req *UpdateEmployeeReq) error {
    res, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("employee data is missing")
        }
        return err
    }

    if req.FullName != nil {
        res.FullName = req.FullName
    }
    if req.Email != nil {
        res.Email = req.Email
    }
    if req.Mobile != nil {
        res.Mobile = req.Mobile
    }
    if req.BirthDate != nil {
        res.BirthDate = req.BirthDate
    }
    if req.Gender != nil {
        res.Gender = req.Gender
    }
    if req.BloodGroup != nil {
        res.BloodGroup = req.BloodGroup
    }
    if req.Designation != nil {
        res.Designation = req.Designation
    }
    if req.EmploymentType != nil {
        res.EmploymentType = req.EmploymentType
    }
    if req.BranchName != nil {
        res.BranchName = req.BranchName
    }
    if req.JoinDate != nil {
        res.JoinDate = req.JoinDate
    }
    if req.Salary != nil {
        res.Salary = req.Salary
    }
    if req.PaymentMethod != nil {
        res.PaymentMethod = req.PaymentMethod
    }
    if req.Nid != nil {
        res.Nid = req.Nid
    }
    if req.Address != nil {
        res.Address = req.Address
    }
    if req.Status != nil {
        res.Status = req.Status
    }
    if req.Image != nil {
        res.Image = req.Image
    }
    if req.Role != nil {
        res.Role = req.Role
    }
    res.UpdatedAt = time.Now()

    if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func (s *employeeService) DeleteEmployee(id uint) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *employeeService) GetAllEmployee(page, pageSize int) ([]Employee, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 10 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    employees, total, err := s.repo.GetAll(pageSize, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("list employees: %w", err)
    }
    return employees, total, nil
}