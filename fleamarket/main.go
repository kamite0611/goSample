package main

import "github.com/gin-gonic/gin"

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 1, Name: "商品2", Price: 2000, Description: "説明1", SoldOut: true},
		{ID: 1, Name: "商品3", Price: 3000, Description: "説明1", SoldOut: false},
	}

	r := gin.Default()
	r.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
