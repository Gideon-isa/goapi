package main

import (
	"github.com/Gideon-isa/productapi/controllers"
	"github.com/Gideon-isa/productapi/models"
	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{Id: 1, Name: "Product 1", Stock: 10, Price: 2.5},
	{Id: 2, Name: "Product 2", Stock: 5, Price: 3.5},
	{Id: 3, Name: "Product 3", Stock: 7, Price: 1.5},
}

func main() {
	r := gin.Default()
	controller := controllers.Init(&products)

	r.GET("/product", controller.ReadProducts)
	r.GET("/product/:id", controller.ReadProductById)
	r.POST("/product", controller.CreateProduct)
	r.PUT("/product/:id", controller.UpdateProduct)
	r.DELETE("/product/:id", controller.DeleteProductById)

	r.Run()

}
