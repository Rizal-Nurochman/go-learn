package initializers

import (
	"log"
	"os"
	models "github.com/NurochmanR/GO-JWT/MODELS"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetToDatabase() {
	var err error
	dsn := os.Getenv("DB")
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database!")
	}

	DB.AutoMigrate(&models.User{})

}