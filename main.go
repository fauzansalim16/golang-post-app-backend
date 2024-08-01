package main

import (
	"backend-go/config"
	"backend-go/routes"
	"log"
	"net/http"
)

func main() {
	err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	r := routes.SetupRoutes()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
