package main

import (
	"fmt"
	"os"
	"rian-anggara/car-rental-api/config"
	"rian-anggara/car-rental-api/docs"
	"rian-anggara/car-rental-api/models"
	"rian-anggara/car-rental-api/routes"
	"rian-anggara/car-rental-api/seed"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			CAR RENTAL SYSTEM
//	@version		0.0.1
//	@description	DOCS CAR RENTAL SYSTEM
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		192.168.1.217:8443
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token

//	@schemes	http https

func main() {
	errl := godotenv.Load("config/.env")
	if errl != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database connection
	db := config.ConnectDB()

	// Run migrations
	errm := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.CarCategory{},
		&models.Location{},
		&models.Car{},
		&models.Booking{},
		&models.Payment{},
		&models.Review{},
		&models.Discount{},
		&models.Insurance{},
		&models.Maintenance{},
	)
	if errm != nil {
		log.Fatalf("Failed to run migrations: %v", errm)
	}

	// Seed default roles
	seed.SeedRoles(db)

	// Set Gin mode based on GIN_MODE env variable
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // fallback
	}
	gin.SetMode(ginMode)

	route := gin.Default()
	route.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			// retrurn response JSON
			c.JSON(200, gin.H{"message": "CAR RENTAL SYSTEM"})
		})
		routes.SetupRouters(v1)
	}
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Host & Port
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	// Print startup info with structured logs
	log.WithFields(log.Fields{
		"component": "main",
		"mode":      gin.Mode(),
		"address":   addr,
	}).Info("🚀 Server is starting")

	if err := route.Run(addr); err != nil {
		log.WithFields(log.Fields{
			"component": "main",
			"error":     err,
		}).Fatal("❌ Failed to start server")
	}
}
