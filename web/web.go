package web

import (
	"fmt"
)

func RunAll() {
	fmt.Println("=== Gin Web Framework ===")
	r := SetupRouter()
	fmt.Println("Router created successfully")

	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET    /ping      - Health check")
	fmt.Println("  GET    /hello     - Query parameter example")
	fmt.Println("  GET    /users     - List all users")
	fmt.Println("  GET    /users/:id - Get user by ID")
	fmt.Println("  POST   /users     - Create user")
	fmt.Println("  PUT    /users/:id - Update user")
	fmt.Println("  DELETE /users/:id - Delete user")
	fmt.Println("  GET    /html      - HTML template")

	_ = r
}
