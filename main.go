package main

import (
	"github.com/gin-gonic/gin"
	"ZCache/routes"
	)
func init(){
	//初始化
}
func main(){
	router := gin.Default()
	router.GET("/:key", routes.Get)
	router.DELETE("/:key", routes.Delete)
	router.POST("/:key/:value", routes.Update)
	router.PUT("/:key/:value", routes.Set)

	router.Run(":8000")
}