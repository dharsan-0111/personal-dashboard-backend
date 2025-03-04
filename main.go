package main

import (
	"fmt"
	"log"
	"personal-dashboard-backend/db"
	"personal-dashboard-backend/routes"
)

func main() {
	fmt.Println("Starting server...")
	db.ConnectDB()
	fmt.Println("Connected to database")

	r := routes.SetUpRoutes()
	log.Println(r, "All Routes")

	port := ":8080"
	log.Println("Server started on port 8080")

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server")
	}
	r.Run(":8080")
}