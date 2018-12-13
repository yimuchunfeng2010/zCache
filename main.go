package main

import (
	"ZCache/global"
	"ZCache/routes"
	"ZCache/routes/mock"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"sync"
	//"os"
	//"os/signal"
)

func init() {
	//初始化
	global.GlobalVar.GRoot = make([]*types.Node, global.Config.MaxLen)
	var i int64
	for i = 0; i < global.Config.MaxLen; i++ {
		global.GlobalVar.GRoot[i] = nil
	}
	global.GlobalVar.GRWLock = new(sync.RWMutex)
	//
	//global.GlobalVar.SigChan = make(chan os.Signal, 1)
	//signal.Notify(global.GlobalVar.SigChan)
}
func main() {
	router := gin.Default()
	v1 := router.Group("/ZCache")
	{
		v1.GET("/:key", routes.Get)
		v1.DELETE("/:key", routes.Delete)
		v1.POST("/:key/:value", routes.Update)
		v1.PUT("/:key/:value", routes.Set)

	}

	v2 := router.Group("/v2")
	{
		v2.GET("/getAll", routes.GetAll)
		v2.GET("/flush", routes.Flush)
		v2.PUT("/import", routes.Import)
	}
	test := router.Group("/mock")
	{
		test.POST("/mockSet", mock.Mock_Set)

	}

	router.Run(":8004")


	//// 启动信号监听
	//go services.SigHandler(global.GlobalVar.SigChan)
	//
	//<-global.GlobalVar.SigChan
}
