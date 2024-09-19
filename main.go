package main

import (
	"github.com/gin-gonic/gin"
)

type NewMessage struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	r.GET("/sender", func(c *gin.Context) {
		var message NewMessage

		if err := c.BindJSON(&message); err != nil {
			return
		}

		Publisher(message.Message)

		c.JSON(200, gin.H{
			"message": "sender complete!",
		})
	})

	r.Run()
}
