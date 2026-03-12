package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "API is running",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Get all users
		v1.GET("/users", getUsers)

		// Get user by ID
		v1.GET("/users/:id", getUserByID)

		// Create new user
		v1.POST("/users", createUser)

		// Update user
		v1.PUT("/users/:id", updateUser)

		// Delete user
		v1.DELETE("/users/:id", deleteUser)
	}

	return r
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate new ID
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    newUser,
	})
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated successfully",
				"data":    updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
