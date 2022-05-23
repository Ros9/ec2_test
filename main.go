package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"os"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/*")
	router.Static("/assets", "./assets")

	router.GET("/upload", func(context *gin.Context) {
		context.HTML(200, "uploader.html", gin.H{})
	})

	router.POST("/attachment", func(context *gin.Context) {
		file, _, err := context.Request.FormFile("attachment_image")
		if err != nil {
			fmt.Println("error while upload file =", err.Error())
			return
		}
		newFile := "assets/" + uuid.New().String()
		out, err := os.Create(newFile)
		if err != nil {
			fmt.Println("error while create file =", err.Error())
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Print("error =", err.Error())
		}
	})

	router.GET("attachment/:name", func(context *gin.Context) {
		fileName := context.Param("name")
		fmt.Println("fileName =", fileName)
		context.HTML(200, "attachment.html", gin.H{
			"fileName": fileName,
		})
	})

	router.Run(":8080")
}
