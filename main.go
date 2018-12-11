package main

import (
	"ZCache/routes"
	"github.com/gin-gonic/gin"
	"ZCache/types"
	"ZCache/global"
)

func init() {
	//初始化
	global.GlobalVar.GRoot = make([]*types.Node, global.Config.MaxLen)
}
func main() {
	router := gin.Default()
	router.GET("/:key", routes.Get)
	router.DELETE("/:key", routes.Delete)
	router.POST("/:key/:value", routes.Update)
	router.PUT("/:key/:value", routes.Set)

	router.Run(":8000")
}
