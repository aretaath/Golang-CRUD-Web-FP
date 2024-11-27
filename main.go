package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/maincontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Main Page
	http.HandleFunc("/", maincontroller.Dashboard)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
