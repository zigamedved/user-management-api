package main

import (
	"log"

	db "github.com/zigamedved/user-management-api/database"
	r "github.com/zigamedved/user-management-api/routers"
)

func main() {
	routes := r.ApiHandleFunctions{}

	log.Printf("Server started")

	db.ConnectDB()
	log.Printf("Database connected")
	defer db.CloseDB()

	router := r.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
