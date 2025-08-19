package main

import (
	"github.com/James-Zeyu-Li/Store_Product/internal/product"
	"github.com/gin-gonic/gin"
)

// ProductModel struct to group product-related handlers
type ProductModel struct{}

func main() {
	product.InitTestData()

	r := gin.Default()
	r.GET("/product/:id", func(c *gin.Context) {
		result, err := product.GetProductById(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, result)
	})
	r.Run(":8080")
}
