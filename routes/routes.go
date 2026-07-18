package routes

import (
	"rent-car-project/controllers"
	"rent-car-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Static file serving
	r.Static("/public", "./public")

	api := r.Group("/api/v1")

	// Public Routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Search route (Public)
	// Categories Routes (Public)
	categories := api.Group("/categories")
	{
		categories.POST("", controllers.CreateCategory)
		categories.GET("", controllers.GetCategories)
		categories.GET("/:id", controllers.GetCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
	}

	cars := api.Group("/cars")
	{
		cars.GET("/search", controllers.SearchCars)
	}

	// Webhook Routes (Public)
	webhooks := api.Group("/webhooks")
	{
		webhooks.POST("/xendit", controllers.XenditWebhook)
	}

	// Owner Protected Routes
	owner := api.Group("/owner")
	owner.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("owner"))
	{
		owner.POST("/cars", controllers.CreateCar)
		owner.GET("/cars", controllers.GetMyCars)
		owner.PUT("/cars/:id", controllers.UpdateCar)
		owner.DELETE("/cars/:id", controllers.DeleteCar)
		owner.POST("/cars/:id/images", controllers.UploadCarImage)
		owner.PUT("/bookings/:id/status", controllers.UpdateBookingStatus)
	}

	// Customer Protected Routes
	customer := api.Group("/customer")
	customer.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("customer"))
	{
		customer.POST("/bookings", controllers.CreateBooking)
		customer.POST("/bookings/:id/reviews", controllers.CreateReview)
	}
}
