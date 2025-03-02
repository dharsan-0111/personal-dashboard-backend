package main

import (
	"fmt"
	// "log"

	"github.com/dharsan-0111/personal-dashboard-backend/startup"
)

func main() {
	startup.Init()
	fmt.Println("Server started")
}