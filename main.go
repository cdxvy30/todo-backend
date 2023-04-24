package main

import (
	"log"

	"todo-backend/routes"
)

func main() {
	r := routes.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
