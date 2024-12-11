package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/maincontroller"
	"go-web/controllers/productcontroller"
	"go-web/controllers/usercontroller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// 1. Main Page
	r.GET("/", maincontroller.Dashboard)

	// 2. Categories
	r.GET("/categories", categorycontroller.Index)
	r.GET("/categories/add", categorycontroller.AddGet)
	r.POST("/categories/add", categorycontroller.AddPost)
	r.GET("/categories/edit", categorycontroller.EditGet)
	r.POST("/categories/edit", categorycontroller.EditPost)
	r.GET("/categories/delete", categorycontroller.Delete)

	// 2. Products
	r.GET("/products", productcontroller.Index)
	r.GET("/products/add", productcontroller.AddGet)
	r.POST("/products/add", productcontroller.AddPost)
	r.GET("/products/detail", productcontroller.Detail)
	r.GET("/products/edit", productcontroller.EditGet)
	r.POST("/products/edit", productcontroller.EditPost)
	r.GET("/products/delete", productcontroller.Delete)

	// 3. Users
	r.GET("/users", usercontroller.Index)
	r.GET("/users/add", usercontroller.AddGet)
	r.POST("/users/add", usercontroller.AddPost)
	r.GET("/users/edit", usercontroller.EditGet)
	r.POST("/users/edit", usercontroller.EditPost)
	r.GET("/users/delete", usercontroller.Delete)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
