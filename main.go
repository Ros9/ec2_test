package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/*")

	router.GET("/upload", func(context *gin.Context) {
		context.HTML(200, "uploader.html", gin.H{})
	})

	router.Run(":8080")
}
