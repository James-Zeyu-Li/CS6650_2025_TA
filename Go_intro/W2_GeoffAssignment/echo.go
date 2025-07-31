package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EchoRequest struct {
	Message string `json:"message"`
}

func postEcho(c *gin.Context) {
	var requestBody EchoRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"echo": requestBody})
}

func main() {
	router := gin.Default()
	router.POST("/echo", postEcho)
	router.Run(":8080")
}
