package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
}

var users = map[int]*User{
	1: {ID: 1, Name: "Alice", Age: 30},
	2: {ID: 2, Name: "Bob", Age: 25},
}
var nextID = 3

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "World")
		c.JSON(200, gin.H{"message": "Hello " + name})
	})

	r.GET("/users", func(c *gin.Context) {
		userList := make([]*User, 0, len(users))
		for _, u := range users {
			userList = append(userList, u)
		}
		c.JSON(200, userList)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, u := range users {
			if string(rune(u.ID+'0')) == id {
				c.JSON(200, u)
				return
			}
		}
		c.JSON(404, gin.H{"error": "User not found"})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		user.ID = nextID
		nextID++
		users[user.ID] = &user
		c.JSON(201, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		for uid, u := range users {
			if string(rune(uid+'0')) == id {
				u.Name = user.Name
				u.Age = user.Age
				c.JSON(200, u)
				return
			}
		}
		c.JSON(404, gin.H{"error": "User not found"})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for uid := range users {
			if string(rune(uid+'0')) == id {
				delete(users, uid)
				c.JSON(200, gin.H{"message": "Deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "User not found"})
	})

	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Hello Gin"})
	})

	return r
}

func GetUser(id int) *User {
	return users[id]
}

func ResetUsers() {
	users = map[int]*User{
		1: {ID: 1, Name: "Alice", Age: 30},
		2: {ID: 2, Name: "Bob", Age: 25},
	}
	nextID = 3
}
