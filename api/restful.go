package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
}

var products = map[int64]*Product{
	1: {ID: 1, Name: "Laptop", Description: "High performance laptop", Price: 999.99, Stock: 10, CreatedAt: time.Now()},
	2: {ID: 2, Name: "Mouse", Description: "Wireless mouse", Price: 29.99, Stock: 50, CreatedAt: time.Now()},
}
var productNextID int64 = 3

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		latency := time.Since(start)
		_ = latency
		_ = method
		_ = path
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer valid-token" {
			c.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Error:   "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}

func SetupAPIRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(LoggerMiddleware())
	r.Use(CORSMiddleware())

	r.GET("/api/v1/health", healthCheck)

	api := r.Group("/api/v1")
	{
		api.Use(AuthMiddleware())

		api.GET("/products", listProducts)
		api.GET("/products/:id", getProduct)
		api.POST("/products", createProduct)
		api.PUT("/products/:id", updateProduct)
		api.DELETE("/products/:id", deleteProduct)

		api.POST("/products/:id/buy", buyProduct)
	}

	return r
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "API is running",
		Data: gin.H{
			"version": "1.0.0",
			"status":  "healthy",
		},
	})
}

func listProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	productList := make([]*Product, 0, len(products))
	for _, p := range products {
		productList = append(productList, p)
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(productList) {
		start = len(productList)
	}
	if end > len(productList) {
		end = len(productList)
	}

	totalPages := (len(productList) + pageSize - 1) / pageSize

	c.JSON(http.StatusOK, PaginatedResponse{
		Success:    true,
		Data:       productList[start:end],
		Page:       page,
		PageSize:   pageSize,
		TotalCount: len(productList),
		TotalPages: totalPages,
	})
}

func getProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	product, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    product,
	})
}

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	product.ID = productNextID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	products[product.ID] = &product
	productNextID++

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	})
}

func updateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	product, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	var updateData Product
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	product.Name = updateData.Name
	product.Description = updateData.Description
	product.Price = updateData.Price
	product.Stock = updateData.Stock
	product.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	if _, exists := products[id]; !exists {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	delete(products, id)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Product deleted successfully",
	})
}

func buyProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	var request struct {
		Quantity int `json:"quantity" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	product, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	if product.Stock < request.Quantity {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Insufficient stock",
		})
		return
	}

	product.Stock -= request.Quantity
	product.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Purchase successful",
		Data: gin.H{
			"product_id": product.ID,
			"quantity":   request.Quantity,
			"total":      product.Price * float64(request.Quantity),
			"remaining":  product.Stock,
		},
	})
}
