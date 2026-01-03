package main

import "github.com/gin-gonic/gin"

func main() {
	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context){
		//map[string]interface{}
		c.JSON(200, gin.H{
			"message" : "AARCS-X API is running!",
			"status" : "success",
		})
	})

	router.Run(":3000")
}