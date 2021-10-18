package main

import (
	"catalog/database"
	"catalog/handler"
)

func main() {

	database.InitDB()

	e := handler.HandlerApi()

	e.Logger.Fatal(e.Start(":8000"))
}
