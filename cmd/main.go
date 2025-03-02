package main

import (
	"fmt"
	"log"

	"github.com/dharsan-0111/personal-dashboard-backend/startup"
)

func main() {
	server := startup.NewServer()

	fmt.Println("Starting server on port 8080...")
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}