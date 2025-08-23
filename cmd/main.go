package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		url := "https://webhook.site/ab01db3b-b40f-4ea4-a98c-33ab41ceaeb7"

		jsonData := []byte(`{"message": "ping"}`)
		response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":81")
}
