package main

import (
	"ZCache/global"
	"ZCache/routes"
	"ZCache/routes/mock"
	"ZCache/types"
	"ZCache/data"
	"github.com/gin-gonic/gin"
	"sync"
)

func init() {
	//初始化
	global.GlobalVar.GRoot = make([]*types.Node, global.Config.MaxLen)
	var i int64
	for i = 0; i < global.Config.MaxLen; i++ {
		global.GlobalVar.GRoot[i] = nil
	}
	global.GlobalVar.GRWLock = new(sync.RWMutex)

	zdata.CoreImport()

}
func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:key", routes.Get)
		v1.DELETE("/:key", routes.Delete)
		v1.POST("/:key/:value", routes.Update)
		v1.PUT("/:key/:value", routes.Set)

	}

	v2 := router.Group("/v2")
	{
		v2.GET("/getAll", routes.GetAll)
		v2.GET("/export", routes.Flush)
		v2.PUT("/import", routes.Import)
		v2.PUT("/deleteAll", routes.DeleteAll)
		v2.PUT("/expension/:size", routes.Expension)
	}

	v3 := router.Group("/v3")
	{
		v3.GET("/getKeyNum",routes.GetKeyNum)
		v3.POST("/incr/:key",routes.Incr)
		v3.POST("/incrBy/:key/:value",routes.IncrBy)
		v3.POST("/decr/:key",routes.Decr)
		v3.POST("/decrBy/:key/:value",routes.DecrBy)
	}

	v4 := router.Group("/v4")
	{
		v4.PUT("/importFromRedis",routes.ImportFromRedis)
		v4.GET("/exportToRedis",routes.ExportToRedis)
	}
	test := router.Group("/mock")
	{
		test.POST("/mockSet", mock.Mock_Set)

	}

	router.Run(":8005")

}
