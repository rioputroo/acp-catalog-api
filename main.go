package main

import (
	"catalog/api"
	productController "catalog/api/product"
	productService "catalog/bussiness/product"
	"catalog/config"
	"catalog/config/rabbitmq"
	"catalog/modules/migration"
	productRepository "catalog/modules/product"
	"fmt"
	"log"
	"os"
	"strconv"

	categoryController "catalog/api/category"
	categoryService "catalog/bussiness/category"
	categoryRepository "catalog/modules/category"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionPostgre() *gorm.DB {

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

func ConnectionMysql() *gorm.DB {

	configDB := map[string]string{

		"DB_Username": config.Config("CATALOG_DB_USERNAME"),
		"DB_Password": config.Config("CATALOG_DB_PASSWORD"),
		"DB_Port":     config.Config("CATALOG_DB_PORT"),
		"DB_Host":     config.Config("CATALOG_DB_ADDRESS"),
		"DB_Name":     config.Config("CATALOG_DB_NAME"),
	}

	fmt.Println(configDB)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	Conn := ConnectionMysql()

	rabbitmq := rabbitmq.RabbitConnection()

	//productRabbitmq := consumer.NewRabbitmqRepository(rabbitmq)

	prodRepository := productRepository.NewProductRepository(Conn, rabbitmq)

	prodRepository.Consume()

	prodService := productService.NewService(prodRepository)
	//prodService := productService.NewService(prodRepository, productRabbitmq)

	prodHandler := productController.NewController(prodService)

	catRepository := categoryRepository.NewCategoryRepository(Conn)
	catService := categoryService.NewService(catRepository)
	catHandler := categoryController.NewController(catService)

	e := echo.New()
	api.HandlerApi(e, prodHandler, catHandler)

	e.Logger.Fatal(e.Start(os.Getenv("CATALOG_APP_PORT")))
}
