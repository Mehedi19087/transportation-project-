package bill

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, productHandler *ProductHandler) {
     v1:= router.Group("/api/v1")
     v1.POST("/products",productHandler.CreateProduct)
     v1.GET("/products/:id",productHandler.GetProduct)
     v1.PUT("/products/:id",productHandler.UpdateProduct)
     v1.DELETE("/products/:id",productHandler.DeleteProduct)
     v1.GET("/products",productHandler.GetAllProducts)

     v1.GET("/companies/products/:companyId", productHandler.GetCompanyProducts)
     
     v1.POST("/bills", productHandler.CreateBill)
     v1.GET("/bills/:id", productHandler.GetBill)
     v1.PUT("/bills/:id", productHandler.UpdateBill)
     v1.DELETE("/bills/:id", productHandler.DeleteBill)
     
     // Changed :productId to :id
     v1.GET("/products/:id/bills", productHandler.GetProductBills)
	 
     v1.PUT("/products/:id/trip-fields", productHandler.UpdateProductTripFields)
     v1.GET("/products/:id/trip-fields", productHandler.GetProductTripFields)
	 v1.PUT("/products/:id/bill-fields", productHandler.UpdateProductBillFields)
     v1.GET("/products/:id/bill-fields", productHandler.GetProductBillFields)
}