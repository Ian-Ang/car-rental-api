package routes

import (
	"rian-anggara/car-rental-api/controllers"
	"rian-anggara/car-rental-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.RouterGroup) {
	// Public Routers
	public := router.Group("/auth")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// Authenticated Routes
	api := router.Group("/")
	api.Use(middlewares.AuthMiddleware())
	{
		// Profile
		api.GET("/users/me", controllers.GetMyProfile)

		// Admin-only User Management
		adminUsers := api.Group("/users")
		adminUsers.Use(middlewares.OnlyAdmin())
		{
			adminUsers.GET("/", controllers.GetAllUsers)
			adminUsers.DELETE("/:id", controllers.DeleteUser)
			adminUsers.GET("/getdeleletuser", controllers.GetDeleteUser)
			adminUsers.POST("/restore/:id", controllers.RestoreUser)
		}

		// Admin-only Roles
		adminRoles := api.Group("/roles")
		adminRoles.Use(middlewares.OnlyAdmin())
		{
			adminRoles.GET("/", controllers.GetRoles)
			adminRoles.POST("/", controllers.CreateRole)
			adminRoles.PUT("/:id", controllers.UpdateRole)
		}

		// 	// Cars (all authenticated users can view, but only admin can create/update)
		// 	api.GET("/cars", controllers.GetCars)
		// 	api.GET("/cars/:id", controllers.GetCarByID)

		// 	adminCars := api.Group("/cars")
		// 	adminCars.Use(middleware.OnlyAdmin())
		// 	{
		// 		adminCars.POST("/", controllers.CreateCar)
		// 		adminCars.PUT("/:id", controllers.UpdateCar)
		// 		adminCars.DELETE("/:id", controllers.DeleteCar)
		// 	}

		// 	// Bookings (customer only)
		// 	customerBookings := api.Group("/bookings")
		// 	customerBookings.Use(middleware.OnlyCustomer())
		// 	{
		// 		customerBookings.POST("/", controllers.CreateBooking)
		// 		customerBookings.GET("/", controllers.GetMyBookings)
		// 	}

		// 	// Admin Bookings view
		// 	adminBookings := api.Group("/admin/bookings")
		// 	adminBookings.Use(middleware.OnlyAdmin())
		// 	{
		// 		adminBookings.GET("/", controllers.GetAllBookings)
		// 	}

		// 	// Reviews (customer only)
		// 	reviews := api.Group("/reviews")
		// 	reviews.Use(middleware.OnlyCustomer())
		// 	{
		// 		reviews.POST("/", controllers.CreateReview)
		// 	}

		// 	// Payments
		// 	api.POST("/payments", controllers.CreatePayment)

		// 	// Discounts (admin only)
		// 	discounts := api.Group("/discounts")
		// 	discounts.Use(middleware.OnlyAdmin())
		// 	{
		// 		discounts.POST("/", controllers.CreateDiscount)
		// 		discounts.GET("/", controllers.GetDiscounts)
		// 	}

		// 	// Insurance (admin only)
		// 	insurance := api.Group("/insurance")
		// 	insurance.Use(middleware.OnlyAdmin())
		// 	{
		// 		insurance.POST("/", controllers.CreateInsurance)
		// 		insurance.GET("/", controllers.GetInsuranceList)
		// 	}

		// 	// Locations (admin only)
		// 	locations := api.Group("/locations")
		// 	locations.Use(middleware.OnlyAdmin())
		// 	{
		// 		locations.POST("/", controllers.CreateLocation)
		// 		locations.GET("/", controllers.GetLocations)
		// 	}

		// 	// Categories (admin only)
		// 	categories := api.Group("/categories")
		// 	categories.Use(middleware.OnlyAdmin())
		// 	{
		// 		categories.POST("/", controllers.CreateCategory)
		// 		categories.GET("/", controllers.GetCategories)
		// 	}

		// 	// Maintenance Records (admin only)
		// 	maintenance := api.Group("/maintenance")
		// 	maintenance.Use(middleware.OnlyAdmin())
		// 	{
		// 		maintenance.POST("/", controllers.CreateMaintenance)
		// 		maintenance.GET("/", controllers.GetMaintenanceList)
		// 	}
	}
}
