package main

import (
	"log"
	"net/http"

	"go-crud/config"
	"go-crud/database"
	"go-crud/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
	database.RunMigration()

	r := routes.ApiRoutes()

	log.Println("Server running on :8000")
	http.ListenAndServe(":8000", r)
}
