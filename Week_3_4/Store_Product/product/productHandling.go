package product

import (
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	productStore = make(map[int64]*ProductModel)
	productMutex = &sync.RWMutex{}
)

// Test Initial Data
func InitTestData() {
	productMutex.Lock()
	defer productMutex.Unlock()

	productStore[1] = &ProductModel{
		ProductId:    1,
		Sku:          "SKU001",
		Manufacturer: "Apple",
		CategoryId:   1,
		Weight:       0.5,
		SomeOtherId:  100,
	}

	productStore[2] = &ProductModel{
		ProductId:    2,
		Sku:          "SKU002",
		Manufacturer: "Samsung",
		CategoryId:   1,
		Weight:       0.6,
		SomeOtherId:  200,
	}
}

func GetProductById(c *gin.Context) (*ProductModel, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) // Convert string ID to int64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid product ID",
			"message": "Invalid ID",
		})
		return nil, errors.New("invalid product ID")
	}

	productMutex.RLock()
	product, ok := productStore[id]
	productMutex.RUnlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Product not found",
			"message": "Product with ID " + idStr + " does not exist",
		})
		return nil, errors.New("product not found")
	}

	c.JSON(http.StatusOK, product)
	return product, nil
}
