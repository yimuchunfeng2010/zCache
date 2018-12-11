package main

import (
	"ZCache/routes"
	"ZCache/routes/mock"
	"github.com/gin-gonic/gin"
	"ZCache/types"
	"ZCache/global"
	"sync"
)

func init() {
	//初始化
	global.GlobalVar.GRoot = make([]*types.Node, global.Config.MaxLen)
	global.GlobalVar.GRWLock = new(sync.RWMutex)
}
func main() {
	router := gin.Default()
	v := router.Group("/ZCache")
	{
		v.GET("/:key", routes.Get)
		v.DELETE("/:key", routes.Delete)
		v.POST("/:key/:value", routes.Update)
		v.PUT("/:key/:value", routes.Set)
	}


	test := router.Group("/mock")
	{
		test.POST("/mockSet", mock.Mock_Set)

	}


	router.Run(":8000")
}
