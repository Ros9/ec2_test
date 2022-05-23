package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/check", func(context *gin.Context) {
		context.String(200, "%s", "Hello, world!")
	})

	router.Run(":8080")
}
