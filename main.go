package main

import (
	"catalog/api"
	productController "catalog/api/product"
	productService "catalog/bussiness/product"
	"catalog/config"
	"catalog/modules/migration"
	productRepository "catalog/modules/product"

	categoryController "catalog/api/category"
	categoryService "catalog/bussiness/category"
	categoryRepository "catalog/modules/category"

	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	p := config.Config("PGPORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Not connected")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("PGHOST"), config.Config("PGUSER"), config.Config("PGPASSWORD"), config.Config("PGDATABASE"), port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection open to database")
	migration.InitMigrate(db)

	return db
}

func main() {

	Conn := Connection()
	prodRepository := productRepository.NewProductRepository(Conn)
	prodService := productService.NewService(prodRepository)
	prodHandler := productController.NewController(prodService)

	catRepository := categoryRepository.NewCategoryRepository(Conn)
	catService := categoryService.NewService(catRepository)
	catHandler := categoryController.NewController(catService)

	e := echo.New()
	api.HandlerApi(e, prodHandler, catHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
