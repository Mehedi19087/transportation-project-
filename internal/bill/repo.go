package bill

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)


type ProductRepo interface {
	 Create(product *Product) error 
	 CreateBill(bill *Bill) error 
	 Get(id uint) (*Product, error)
	 Update(product *Product) error 
	 Delete(id uint) error 
	 GetAll(offset , limit int) ([]Product, int64 ,error)
	 GetByCompany(companyID uint, offset , limit int) ([]Product, int64, error)

	 GetBill(id uint) (*Bill, error)
	 GetBillsByProduct(productID uint, offset , limit int) ([]Bill, int64, error)
	 UpdateBill(bill *Bill) error
	 DeleteBill(id uint) error
	 UpdateProductTripFields(productID uint, tripFields pq.StringArray) error
	 GetProductTripFields(productID uint) (pq.StringArray, error)

	 UpdateProductBillFields(productID uint, billFields pq.StringArray) error
     GetProductBillFields(productID uint) (pq.StringArray, error) 
}

type productRepo struct {
	 db *gorm.DB 
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	 return &productRepo {db : db}
}

func(r *productRepo) Create(product *Product) error {
	 return r.db.Create(product).Error 
}

func(r *productRepo) Get(id uint) (*Product, error) {
	 var res Product 
	 if err := r.db.First(&res, id).Error; err != nil{
		 return nil, err  
	 }
	 return &res , nil 
}

func(r *productRepo) Update(product *Product) error {
	 return r.db.Save(product).Error
}

func(r *productRepo) Delete(id uint) error {
	 return r.db.Delete(&Product{}, id).Error 
}

func(r *productRepo) GetAll(offset, limit int) ([]Product, int64 , error) {
	   var res int64 
	   var products []Product
	   if err:= r.db.Model(&Product{}).Count(&res).Error; err != nil {
		  return nil, 0, err 
	   }
	   if err:= r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		    return nil , 0 , err
	   }
	   return products, res, nil 
}


func(r *productRepo) GetByCompany(companyID uint, offset , limit int) ([]Product, int64, error) {
	   var total int64 
	   var products []Product 

	   if err := r.db.Model(&Product{}).Where("company_id = ?", companyID).Count(&total).Error ; err != nil {
		   return nil , 0, err 
	   }

	   if err := r.db.Where("company_id = ?", companyID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		    return nil, 0, err 
	   }
	   return products, total, nil 
	   
}

func(r *productRepo) CreateBill(bill *Bill) error {
	 return r.db.Create(bill).Error 
}

func(r *productRepo) GetBill(id uint) (*Bill, error) {
    var bill Bill 
    if err := r.db.First(&bill, id).Error; err != nil {
        return nil, err  
    }
    return &bill, nil 
}

func(r *productRepo) GetBillsByProduct(productID uint, offset , limit int) ([]Bill, int64, error) {
	 var total int64 
	 var bills []Bill
	 
	 if err := r.db.Model(&Bill{}).Where("product_id = ?",productID).Count(&total).Error; err != nil {
		 return nil, 0, err 
	 }

	 if err := r.db.Model(&Bill{}).Where("product_id = ?",productID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&bills).Error; err != nil {
		 return nil, 0, err 
	 }
	 return bills , total, nil 
}

func(r *productRepo) UpdateBill(bill *Bill) error {
    return r.db.Save(bill).Error
}

func(r *productRepo) DeleteBill(id uint) error {
    return r.db.Delete(&Bill{}, id).Error 
}

func(r *productRepo) UpdateProductTripFields(productID uint, tripFields pq.StringArray) error {
    return r.db.Model(&Product{}).Where("id = ?", productID).Update("trip_fields", tripFields).Error
}

func(r *productRepo) GetProductTripFields(productID uint) (pq.StringArray, error) {
	 var product Product 

	 if err := r.db.Select("trip_fields").First(&product, productID).Error ; err != nil {
		  return nil , err 
	 }
	 return product.TripFields, nil 
}

func(r *productRepo) UpdateProductBillFields(productID uint, billFields pq.StringArray) error {
    return r.db.Model(&Product{}).Where("id = ?", productID).Update("bill_fields", billFields).Error
}

func(r *productRepo) GetProductBillFields(productID uint) (pq.StringArray, error) {
     var product Product 
     if err := r.db.Select("bill_fields").First(&product, productID).Error ; err != nil {
          return nil , err 
     }
     return product.BillFields, nil 
}