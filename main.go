package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yujen77300/URL-Shortener-Gin/controllers/handlers"
	_ "github.com/yujen77300/URL-Shortener-Gin/models"
)



func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	port := ":8080"

	router.Static("/css", "./static/css")
	router.Static("/javascript", "./static/javascript")
	router.Static("/image", "./static/image")

	router.LoadHTMLGlob("views/*.html")


	router.GET("/", handlers.Home)
	router.POST("/", handlers.Shorten)
	router.GET("/:shortenString", handlers.Redirect)


	router.Run(port)

}
