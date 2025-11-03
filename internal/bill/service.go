package bill

import (
	"errors"
	"time"
	"fmt"
	"transportation/internal/customer"

	"gorm.io/gorm"
    "github.com/lib/pq"
)

type ProductService interface {
	CreateProduct(req *CreateProductReq) error
	GetProduct(id uint) (*Product, error)
	UpdateProduct(id uint, req *UpdateProductReq) error 
	DeleteProduct(id uint) error 
	GetAllProducts(page, pageSize int) ([]Product,int64, error)

	GetCompanyProducts(companyID uint, page, pageSize int) ([]Product, int64, error)

	CreateBill(req *CreateBillReq) error
	GetBill(id uint) (*Bill, error)
    GetProductBills(productID uint, page, pageSize int) ([]Bill, int64, error)
    UpdateBill(id uint, req *UpdateBillReq) error
    DeleteBill(id uint) error
    UpdateProductTripFields(productID uint, req *UpdateProductTripFieldsReq) error
    GetProductTripFields(productID uint) ([]string, error)

    UpdateProductBillFields(productID uint, req *UpdateProductBillFieldsReq) error  // Add this line
    GetProductBillFields(productID uint) ([]string, error)
}

type productService struct {
	repo ProductRepo
	customerRepo customer.CustomerRepo
}

func NewProductService(repo ProductRepo, customerRepo customer.CustomerRepo) ProductService {
	return &productService{
		repo: repo,
		customerRepo: customerRepo,
	}
}

func (s *productService) CreateProduct(req *CreateProductReq) error {
	if req.Name == "" {
		return errors.New("product name is required")
	}
	if req.CompanyID == 0 {
        return errors.New("company_id is required")
    }
    _, err := s.customerRepo.Get(req.CompanyID)
	if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("customer not found")
        }
        return fmt.Errorf("failed to validate company: %w", err)
	}

	product:= &Product{
		Name: req.Name,
		CompanyID: req.CompanyID,
		Alt: req.Alt,
		Vat: req.Vat,
		CreatedAt: time.Now().UTC(),
	}
	err = s.repo.Create(product)
	if err != nil {
		 return err 
	}
	return nil 
}

func(s *productService) GetProduct(id uint) (*Product, error) {
	  res, err := s.repo.Get(id)
	  if err != nil {
		 return nil, err 
	  }
	  return res, nil 
}

func(s *productService) UpdateProduct(id uint, req *UpdateProductReq) error {
	 res, err := s.repo.Get(id) 
	 if err != nil {
		  if err == gorm.ErrRecordNotFound {
			 return errors.New("product data does not exist")
		  }
		  return err 
	 }
	 if req.Name != "" {
		 res.Name = req.Name
	 }
	 if req.Alt!= nil {
		 res.Alt= req.Alt
	 }
	 if req.Vat!= nil{
		 res.Vat= req.Vat
	 }
	 if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func(s *productService) DeleteProduct(id uint) error {
	 err:=s.repo.Delete(id)
	 if err != nil {
		 return  err 
	 }
	 return nil 
}

func(s *productService) GetAllProducts(page, pageSize int) ([]Product, int64,error) {
	 if page < 1 {
		 page =1 
	 }
	 if pageSize < 1 || pageSize>10 {
		 pageSize=10 
	 }
	 offset:= (page-1)*pageSize

	 products, total, err := s.repo.GetAll(offset, pageSize)

	 if err != nil {
        return nil, 0, err
    }
    return products, total, nil
}

func (s *productService) GetCompanyProducts(companyID uint, page, pageSize int) ([]Product, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize
    
    items, total, err := s.repo.GetByCompany(companyID, offset, pageSize)
    if err != nil {
        return nil, 0, fmt.Errorf("list products by company: %w", err)
    }
    return items, total, nil
}

func (s *productService) CreateBill(req *CreateBillReq) error {
    if req.ProductID == 0 {
        return errors.New("product_id is required")
    }

    // Verify product exists
    _, err := s.repo.Get(req.ProductID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("product not found")
        }
        return fmt.Errorf("failed to verify product: %w", err)
    }
	bill := &Bill{
        Category:              req.Category,
        ProductID:             req.ProductID,
        VehicleNo:             req.VehicleNo,
        CustomerName:          req.CustomerName,
        ChallanNo:             req.ChallanNo,
        DistributorName:       req.DistributorName,
        DealerName:            req.DealerName,
        DriverName:            req.DriverName,
        FromLocation:          req.FromLocation,
        Destination:           req.Destination,
        Product:               req.Product,
        Portfolio:             req.Portfolio,
        Goods:                 req.Goods,
        Quantity:              req.Quantity,
        BikeQty:               req.BikeQty,
        VehicleSize:           req.VehicleSize,
        Status:                req.Status,
        UnloadCharge:          req.UnloadCharge,
        VehicleRentWithVATTax: req.VehicleRentWithVatTax,
        VehicleRent:           req.VehicleRent,
        Dropping:              req.Dropping,
        Alt5:                  req.Alt5,
        Vat10:                 req.Vat10,
        TotalRate:             req.TotalRate,
        Advance:               req.Advance,
		Due:                   req.Due,
        Total:                 req.Total,
        Profit:                req.Profit,
        BodyFare:              req.BodyFare,
        FuelCost:              req.FuelCost,
        Amount:                req.Amount,
        TotalAmount:           req.TotalAmount,
        DoNumber:              req.DoNumber,
        CoNumber:              req.CoNumber,
        CreatedAt:             time.Now().UTC(),
    }

    if err := s.repo.CreateBill(bill); err != nil {
        return fmt.Errorf("failed to create bill: %w", err)
    }

    return nil
}

func (s *productService) GetBill(id uint) (*Bill, error) {
    if id == 0 {
        return nil, errors.New("invalid bill id")
    }
    
    bill, err := s.repo.GetBill(id) 
    if err != nil {
        return nil, err 
    }
    return bill, nil 
}

func (s *productService) GetProductBills(productID uint, page, pageSize int) ([]Bill, int64, error) {
    if productID == 0 {
        return nil, 0, errors.New("product_id is required")
    }

    // Verify product exists
    _, err := s.repo.Get(productID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, 0, errors.New("product not found")
        }
        return nil, 0, fmt.Errorf("failed to verify product: %w", err)
    }

    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    bills, total, err := s.repo.GetBillsByProduct(productID, offset, pageSize)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to get bills by product: %w", err)
    }

    return bills, total, nil
}


// ...existing code...

func(s *productService) UpdateBill(id uint, req *UpdateBillReq) error {
    bill, err := s.repo.GetBill(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("bill data does not exist")
        }
        return err
    }

    // If ProductID is being updated, verify the new product exists
    if req.ProductID != nil && *req.ProductID != 0 {
        _, err := s.repo.Get(*req.ProductID)
        if err != nil {
            if err == gorm.ErrRecordNotFound {
                return errors.New("product not found")
            }
            return fmt.Errorf("failed to verify product: %w", err)
        }
        bill.ProductID = *req.ProductID
    }

    if req.Category != nil {
        bill.Category = req.Category
    }
    if req.VehicleNo != nil {
        bill.VehicleNo = req.VehicleNo
    }
    if req.CustomerName != nil {
        bill.CustomerName = req.CustomerName
    }
    if req.ChallanNo != nil {
        bill.ChallanNo = req.ChallanNo
    }
    if req.DistributorName != nil {
        bill.DistributorName = req.DistributorName
    }
    if req.DealerName != nil {
        bill.DealerName = req.DealerName
    }
    if req.DriverName != nil {
        bill.DriverName = req.DriverName
    }
    if req.FromLocation != nil {
        bill.FromLocation = req.FromLocation
    }
    if req.Destination != nil {
        bill.Destination = req.Destination
    }
    if req.Product != nil {
        bill.Product = req.Product
    }
    if req.Portfolio != nil {
        bill.Portfolio = req.Portfolio
    }
    if req.Goods != nil {
        bill.Goods = req.Goods
    }
    if req.Quantity != nil {
        bill.Quantity = req.Quantity
    }
    if req.BikeQty != nil {
        bill.BikeQty = req.BikeQty
    }
    if req.VehicleSize != nil {
        bill.VehicleSize = req.VehicleSize
    }
    if req.Status != nil {
        bill.Status = req.Status
    }
    if req.UnloadCharge != nil {
        bill.UnloadCharge = req.UnloadCharge
    }
    if req.VehicleRentWithVatTax != nil {
        bill.VehicleRentWithVATTax = req.VehicleRentWithVatTax
    }
    if req.VehicleRent != nil {
        bill.VehicleRent = req.VehicleRent
    }
    if req.Dropping != nil {
        bill.Dropping = req.Dropping
    }
    if req.Alt5 != nil {
        bill.Alt5 = req.Alt5
    }
    if req.Vat10 != nil {
        bill.Vat10 = req.Vat10
    }
    if req.TotalRate != nil {
        bill.TotalRate = req.TotalRate
    }
    if req.Advance != nil {
        bill.Advance = req.Advance
    }
    if req.Due != nil {
        bill.Due = req.Due
    }
    if req.Total != nil {
        bill.Total = req.Total
    }
    if req.Profit != nil {
        bill.Profit = req.Profit
    }
    if req.BodyFare != nil {
        bill.BodyFare = req.BodyFare
    }
    if req.FuelCost != nil {
        bill.FuelCost = req.FuelCost
    }
    if req.Amount != nil {
        bill.Amount = req.Amount
    }
    if req.TotalAmount != nil {
        bill.TotalAmount = req.TotalAmount
    }
    if req.DoNumber != nil {
        bill.DoNumber = req.DoNumber
    }
    if req.CoNumber != nil {
        bill.CoNumber = req.CoNumber
    }

    if err := s.repo.UpdateBill(bill); err != nil {
        return err
    }
    return nil
}

func(s *productService) DeleteBill(id uint) error {
    if id == 0 {
        return errors.New("invalid bill id")
    }

    // Check if bill exists before deleting
    _, err := s.repo.GetBill(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("bill not found")
        }
        return fmt.Errorf("failed to verify bill: %w", err)
    }

    if err := s.repo.DeleteBill(id); err != nil {
        return fmt.Errorf("failed to delete bill: %w", err)
    }

    return nil
}

func(s *productService) UpdateProductTripFields(productID uint, req *UpdateProductTripFieldsReq) error {
    if productID == 0 {
        return errors.New("invalid product id")
    }

    // Check if product exists
    _, err := s.repo.Get(productID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("product not found")
        }
        return fmt.Errorf("failed to verify product: %w", err)
    }


    if len(req.TripFields) == 0 {
        return errors.New("trip_fields cannot be empty")
    }
    tripFields := pq.StringArray(req.TripFields)

    if err := s.repo.UpdateProductTripFields(productID, tripFields); err != nil {
        return fmt.Errorf("failed to update trip fields: %w", err)
    }

    return nil
}

func(s *productService) GetProductTripFields(productID uint) ([]string, error) {
    if productID == 0 {
        return nil, errors.New("invalid product id")
    }

    tripFields, err := s.repo.GetProductTripFields(productID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("product not found")
        }
        return nil, fmt.Errorf("failed to get trip fields: %w", err)
    }

    return []string(tripFields), nil
}


func(s *productService) UpdateProductBillFields(productID uint, req *UpdateProductBillFieldsReq) error {
    if productID == 0 {
        return errors.New("invalid product id")
    }

    // Check if product exists
    _, err := s.repo.Get(productID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("product not found")
        }
        return fmt.Errorf("failed to verify product: %w", err)
    }

    if len(req.BillFields) == 0 {
        return errors.New("bill_fields cannot be empty")
    }

    // Convert []string to pq.StringArray
    billFields := pq.StringArray(req.BillFields)
    
    if err := s.repo.UpdateProductBillFields(productID, billFields); err != nil {
        return fmt.Errorf("failed to update bill fields: %w", err)
    }

    return nil
}

func(s *productService) GetProductBillFields(productID uint) ([]string, error) {
    if productID == 0 {
        return nil, errors.New("invalid product id")
    }

    billFields, err := s.repo.GetProductBillFields(productID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("product not found")
        }
        return nil, fmt.Errorf("failed to get bill fields: %w", err)
    }

    // Convert pq.StringArray back to []string
    return []string(billFields), nil
}