package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Product 结构体保持不变
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 内存数据库和互斥锁也保持不变，因为这部分逻辑与Web框架无关
var productStore = make(map[string]Product)
var mu sync.RWMutex

// getProductHandlerGin 是使用 Gin 的 get 处理器
// 注意函数签名变成了 func(c *gin.Context)
func getProductHandlerGin(c *gin.Context) {
	// Gin 提供了更简洁的方式来获取 URL 参数
	id := c.Param("id")

	// 互斥锁逻辑不变
	mu.RLock()
	product, ok := productStore[id]
	mu.RUnlock()

	if !ok {
		// Gin 使用 c.JSON 来返回 JSON 响应，并可以方便地设置状态码
		// gin.H 是 map[string]interface{} 的一个快捷方式
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// 成功时返回 200 OK 和产品数据
	c.JSON(http.StatusOK, product)
}

// createProductHandlerGin 是使用 Gin 的 post 处理器
func createProductHandlerGin(c *gin.Context) {
	var newProduct Product

	// Gin 提供了强大的请求体绑定功能，c.ShouldBindJSON 会自动解析 JSON
	// 并将其填充到 newProduct 结构体中。
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 互斥锁逻辑不变
	mu.Lock()
	productStore[newProduct.ID] = newProduct
	mu.Unlock()

	// 成功时返回 201 Created 和新创建的产品数据
	c.JSON(http.StatusCreated, newProduct)
}

func main() {
	// 添加初始数据
	productStore["apple-iphone-15"] = Product{ID: "apple-iphone-15", Name: "iPhone 15 Pro", Price: 1099.00}
	productStore["samsung-galaxy-s24"] = Product{ID: "samsung-galaxy-s24", Name: "Galaxy S24 Ultra", Price: 1299.00}
	productStore["google-pixel-8"] = Product{ID: "google-pixel-8", Name: "Google Pixel 8", Price: 699.00}
	productStore["sony-wh-1000xm5"] = Product{ID: "sony-wh-1000xm5", Name: "Sony Headphones", Price: 399.00}

	// 创建一个 Gin 路由器
	// gin.Default() 会返回一个带有 Logger 和 Recovery 中间件的引擎
	router := gin.Default()

	// 定义路由规则，Gin 的风格更简洁
	// 注意 URL 参数的语法从 {id} 变成了 :id
	router.GET("/products/:id", getProductHandlerGin)
	router.POST("/products", createProductHandlerGin)

	log.Println("Server starting on port 8080 with Gin")

	// 启动服务器，Gin 的 Run 方法封装了 http.ListenAndServe
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
