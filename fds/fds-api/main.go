package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	// Set up database connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Disable caching
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Register API endpoints
	r.GET("/api/address/totals", getAddressTotals)
	r.GET("/api/suspicious", getSuspiciousTransactions)
	r.GET("/api/blacklist", getBlacklist)
	r.GET("/api/address/related", getRelatedAddresses)
	r.GET("/api/address/transactions", getTransactionsByAddress)
	r.GET("/api/suspicious/related", getRelatedTransactionsOfSuspicious)
	r.GET("/api/balance/eth", getETHBalance)
	r.GET("/api/balance/evnd", getEVNDBalance)
	r.POST("/api/blacklist", blacklistAddresses)
	r.POST("/api/unblacklist", unblacklistAddresses)
	r.DELETE("/api/blacklist", deleteBlacklistAddress)
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("Starting API server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
