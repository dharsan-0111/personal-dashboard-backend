package main

import (
	"fmt"
	// "log"

	"personal-dashboard-backend/startup"
)

func main() {
	startup.Init()
	fmt.Println("Server started")
}