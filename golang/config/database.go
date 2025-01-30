package config

import (
	"fmt"
	"log"
	"os"

	"tech-test/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	user, password, host, port, dbName)
		

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = DB.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}


	fmt.Println("Database connection established")
}
