package main

import (
	"log"
	employee "transportation/internal/Employee"
	outsidetrip "transportation/internal/OutSide_Trip"
	trip "transportation/internal/Trip"
	"transportation/internal/bill"
	"transportation/internal/customer"
	"transportation/internal/database"
	"transportation/internal/dealer"
	"transportation/internal/driver"
	"transportation/internal/media"
	routepricing "transportation/internal/routePricing"
	"transportation/internal/vehicle"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		 log.Panic("Database Initialization Issue")
	}
	gin.SetMode(gin.ReleaseMode)

    uploader, err := media.NewCloudinaryUploader()
    if err != nil { 
		log.Println("cloudinary initialization failed")
	}


	customerRepo := customer.NewCustomerRepo(db)
	dealerRepo   := dealer.NewDealerRepo(db)
	driverRepo   :=driver.NewDriverRepo(db)
	vehicleRepo := vehicle.NewVehicleRepo(db)
	routePricingRepo := routepricing.NewRoutePricingRepo(db)
	productRepo :=bill.NewProductRepo(db)
	tripRepo:=    trip.NewRepository(db)
	employeeRepo:= employee.NewEmployeeRepo(db)
	outsideRepo := outsidetrip.NewOutSideTripRepo(db)
    
   


	customerService:=customer.NewCustomerService(customerRepo)
	dealerService  :=dealer.NewDealerService(dealerRepo)
	driverService  :=driver.NewDriverService(driverRepo)
	vehicleService := vehicle.NewVehicleService(vehicleRepo)
	routePricingService := routepricing.NewRoutePricingService(routePricingRepo)
	productService:= bill.NewProductService(productRepo,customerRepo)
	tripService:=    trip.NewService(tripRepo)
	employeeService:= employee.NewEmployeeService(employeeRepo)
	outsideService:= outsidetrip.NewOutSideTripService(outsideRepo)
    


	customerHanlder:=customer.NewCustomerHandler(customerService)
	dealerHandler  :=dealer.NewDealerHandler(dealerService)
	driverHandler:=  driver.NewDriverHandler(driverService,uploader)
	vehicleHandler := vehicle.NewVehicleHandler(vehicleService,uploader)
	routePricingHandler := routepricing.NewRoutePricingHandler(routePricingService)
	productHandler:= bill.NewProductHandler(productService)
	tripHandler:=    trip.NewHandler(tripService)
	employeeHandler:= employee.NewEmployeeHandler(employeeService)
	outsideHandler    := outsidetrip.NewOutSideTripHandler(outsideService)
    



	router:= gin.Default()
	router.MaxMultipartMemory = 8 << 20

	

    router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	customer.SetupRoutes(router,customerHanlder)
	dealer.SetupRoutes(router,dealerHandler)
	driver.SetupRoutes(router,driverHandler)
	vehicle.SetupRoutes(router, vehicleHandler)
	routepricing.SetupRoutes(router, routePricingHandler)
	bill.SetupRoutes(router, productHandler)
	trip.SetupRoutes(router,tripHandler)
	employee.SetupRoutes(router,employeeHandler)
	outsidetrip.SetupRoutes(router,outsideHandler)


	if err := router.Run(":8080"); err != nil {
    log.Fatalf("server failed: %v", err)
    }
    
}