package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/:wordToSearch", func(c *gin.Context) {
		wordToSearch := c.Param("wordToSearch")
		outputJSON := getMeaning(wordToSearch)
		c.JSON(http.StatusOK, outputJSON)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getMeaning(wordToSearch string) interface{} {
	return Search(wordToSearch)
}
