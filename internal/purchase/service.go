package purchase

import (
	"time"
)

type PurchaseService interface {
    CreatePurchase(req *CreatePurchase) error
    UpdatePurchase(id uint, req *UpdatePurchase) error
    GetPurchase(id uint) (*Purchase, error)
    DeletePurchase(id uint) error
    GetAllPurchase(page, pageSize int) ([]Purchase, int64, error)
}

type purchaseService struct {
    repo PurchaseRepo
}

func NewPurchaseService(repo PurchaseRepo) PurchaseService {
    return &purchaseService{repo: repo}
}

func (s *purchaseService) CreatePurchase(req *CreatePurchase) error {
    // map DTO -> model (model uses pointers)
    p := &Purchase{
        SupplierName:  strPtr(req.SupplierName),
        OrderDate:     timePtr(req.OrderDate),
        DeliveryDate:  timePtr(req.DeliveryDate),
        Items:         strPtr(req.Items),
        Quantity:      strPtr(req.Quantity),
        Category:      strPtr(req.Category),
        TotalAmount:   floatPtr(req.TotalAmount),
        VehicleNo:     strPtr(req.VehicleNo),
        Notes:         req.Notes,
        PaymentStatus: strPtr(req.PaymentStatus),
        BillImage:     req.BillImage,
    }
    return s.repo.Create(p)
}

func (s *purchaseService) UpdatePurchase(id uint, req *UpdatePurchase) error {
    existing, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    // overwrite like driver Update flow
    existing.SupplierName = strPtr(req.SupplierName)
    existing.OrderDate = timePtr(req.OrderDate)
    existing.DeliveryDate = timePtr(req.DeliveryDate)
    existing.Items = strPtr(req.Items)
    existing.Quantity = strPtr(req.Quantity)
    existing.Category = strPtr(req.Category)
    existing.TotalAmount = floatPtr(req.TotalAmount)
    existing.VehicleNo = strPtr(req.VehicleNo)
    existing.Notes = req.Notes
    existing.PaymentStatus = strPtr(req.PaymentStatus)
    existing.BillImage = req.BillImage

    return s.repo.Update(existing)
}

func (s *purchaseService) GetPurchase(id uint) (*Purchase, error) {
    return s.repo.Get(id)
}

func (s *purchaseService) DeletePurchase(id uint) error {
    return s.repo.Delete(id)
}

func (s *purchaseService) GetAllPurchase(page, pageSize int) ([]Purchase, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize
    return s.repo.GetAll(offset, pageSize)
}

// helpers
func strPtr(s string) *string { return &s }
func floatPtr(f float64) *float64 { return &f }
func timePtr(t any) *time.Time {
    switch v := t.(type) {
    case time.Time:
        if v.IsZero() {
            return nil
        }
        return &v
    default:
        return nil
    }
}