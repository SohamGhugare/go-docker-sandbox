package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handling the root route
func HandleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "root",
	})
}

func main(){
	// Initializing server
	r := gin.Default()

	// Registering routes
	r.GET("/", HandleRoot)

	// Running the instance
	r.Run()
}