package main

import (
	"catalog/config"
	"catalog/config/rabbitmq"
	"catalog/modules/migration"
	productRepository "catalog/modules/product"
	"fmt"
	"log"
	"strconv"

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

	prodRepository := productRepository.NewProductRepository(Conn, rabbitmq)

	prodRepository.Consume(Conn)
}
