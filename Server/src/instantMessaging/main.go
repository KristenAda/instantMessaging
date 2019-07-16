package main

import (
	"github.com/gin-gonic/gin"
	"instantMessaging/web"
	"log"
	"net/http"
)

func main() {
	webCon := gin.Default()
	webCon.GET("/user",userHello)

	go web.LongLink()
	err := http.ListenAndServe(":8888",webCon)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func userHello(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"hello world!",
	})
}
