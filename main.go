package main

import (
	"bookmyshow/postgresql"
	"bookmyshow/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	postgresql.ConnectDB()
	r := gin.Default()
	routes.Router(r)
	r.Run(":8080")
}
