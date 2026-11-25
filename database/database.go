package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"api/models"
)

var DB *gorm.DB

func Connect() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database: ", err)
	}

	DB = db

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("could not migrate the database: ", err)
	}

	log.Println("database connected")
}
