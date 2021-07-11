package main

import (
	"log"
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
	r.POST("/api/tg", func(c *gin.Context) {
		var body Update
		err := c.ShouldBindJSON(&body)
		if err != nil {
			// error handle
		}

		log.Println(body.Message.Text)
		outputJSON := getMeaning(body.Message.Text)
		SendTextToTelegramChat(body.Message.Chat.Id, PreprocessingJSONToString(outputJSON))

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getMeaning(wordToSearch string) wordMeaning {
	return Search(wordToSearch)
}
