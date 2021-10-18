package database

import (
	"catalog/config"
	"catalog/models"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	p := config.Config("PGPORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("PGHOST"), config.Config("PGUSER"), config.Config("PGPASSWORD"), config.Config("PGDATABASE"), port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection open to database")
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Category{})
}
