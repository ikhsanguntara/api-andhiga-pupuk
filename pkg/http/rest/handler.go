package rest

import (
	"food_delivery_api/pkg/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Handler(s service.Service) *gin.Engine {
	r := gin.Default()
	setupCORS(r)

	// Public API
	r.GET("/health", getHealthStatus)
	r.POST("/api/v1/login", login(s))

	// Protected API
	v1 := r.Group("/api/v1")
	// v1.Use(middleware.JWT())
	{
		// Users
		v1.POST("/users", addUser(s))
		v1.GET("/users", getUsers(s))
		v1.GET("/users/:id", getUser(s))
		v1.PUT("/users/:id", editUser(s))
		v1.DELETE("/users/:id", removeUser(s))

		// Products
		v1.POST("/products", addProduct(s))
		v1.GET("/products", getProducts(s))
		v1.GET("/products/:id", getProduct(s))
		v1.PUT("/products/:id", editProduct(s))
		v1.DELETE("/products/:id", removeProduct(s))

	}

	return r
}

func setupCORS(r *gin.Engine) {
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, PATCH, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          1 * time.Minute,
		Credentials:     false,
		ValidateHeaders: false,
	}))
}

func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
