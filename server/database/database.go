package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DelaRicch/klock-ecommerce/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database \n %v", err)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	if os.Getenv("APP_ENV") == "development" {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
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
