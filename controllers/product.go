package controllers

import (
	"net/http"
	"strconv"

	"github.com/Gideon-isa/productapi/models"
	"github.com/gin-gonic/gin"
)

type ProductRepo struct {
	Products *[]models.Product
}

// Initializes a ProductRepo
func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}

// CRUD operations for controllers
// CreateProducts binds the struct into json and serves it
func (repo *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product

	c.BindJSON(&product)
	err := models.CreateProduct(repo.Products, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (repo *ProductRepo) ReadProducts(c *gin.Context) {
	c.JSON(http.StatusOK, repo.Products)
}

func (repo *ProductRepo) ReadProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	product := models.ReadProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}

func (repo *ProductRepo) UpdateProduct(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	var product models.Product
	//this binds the json values sent as request to the product variable
	c.BindJSON(&product)

	if product.Id != int(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	updatedProduct := models.UpdatedProductById(repo.Products, &product)
	c.JSON(http.StatusOK, &updatedProduct)
}

func (repo *ProductRepo) DeleteProductById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)

	product := models.DeleteProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}
