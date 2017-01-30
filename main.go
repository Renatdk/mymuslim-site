package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main(){
	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	r.GET("/index", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",})
	})
	r.Run(":8081")
}
