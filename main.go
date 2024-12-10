package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/maincontroller"
	"go-web/controllers/productcontroller"
	"go-web/controllers/usercontroller"
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
	
	// 2. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)
	
	// 3. Users
	http.HandleFunc("/users", usercontroller.Index)
	http.HandleFunc("/users/add", usercontroller.Add)
	http.HandleFunc("/users/edit", usercontroller.Edit)
	http.HandleFunc("/users/delete", usercontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
