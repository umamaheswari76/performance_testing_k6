package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//10lac req with 100 thread


func main() {
	r := gin.Default()
	r.POST("/k6test", func(c *gin.Context) {
		// Parse the JSON request body
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Modify the JSON data
		// requestBody["message"] = "Modified: " + requestBody["message"].(string)

		// Send the modified JSON as a response
		c.JSON(http.StatusOK, requestBody)
	})

	r.Run(":8080")
}

