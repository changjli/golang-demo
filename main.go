package main

import (
	"user_service/routes"
)

func main() {
	// fmt.Println("hello world");

	router := routes.SetupRoutes()

	router.Run(":8015")
}
