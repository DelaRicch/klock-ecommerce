package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DelaRicch/klock-ecommerce/klock-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDb() {
	// Load environment variables from .env file 
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database \n %v", err)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	if os.Getenv("APP_ENV") == "development" {
		// Auto Migrate User model
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("Failed to migrate User model: %v", err)
		}

		// Auto Migrate Product model
		if err := db.AutoMigrate(&models.Product{}); err != nil {
			log.Fatalf("Failed to migrate Product model: %v", err)
		}

		// Auto Migrate ProductGalleryImage model
		if err := db.AutoMigrate(&models.ProductGalleryImage{}); err != nil {
			log.Fatalf("Failed to migrate ProductGalleryImage model: %v", err)
		}

	}

	DB = db
}

func CloseDb() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}
	sqlDB.Close()
}
