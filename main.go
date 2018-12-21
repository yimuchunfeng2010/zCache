package main

import (
	"ZCache/global"
	"ZCache/routes"
	"ZCache/routes/mock"
	"ZCache/services"
	"ZCache/types"
	"ZCache/data"
	"ZCache/task"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	//初始化
	global.GlobalVar.GRoot = make([]*types.Node, global.Config.MaxLen)
	var i int64
	for i = 0; i < global.Config.MaxLen; i++ {
		global.GlobalVar.GRoot[i] = nil
	}

	zdata.CoreImport()

	// 初始化日志
	global.GlobalVar.GLogInfoHead = new(types.LogInfoNode)
	global.GlobalVar.GLogInfoTail = global.GlobalVar.GLogInfoHead
	global.GlobalVar.GLogWarningHead = new(types.LogInfoNode)
	global.GlobalVar.GLogWarningTail = global.GlobalVar.GLogWarningHead
	global.GlobalVar.GLogErrorHead = new(types.LogInfoNode)
	global.GlobalVar.GLogErrorTail = global.GlobalVar.GLogErrorHead

	// 初始化zk
	services.ZookeeperInit()
	// 注册zk集群节点
	err := services.RegisterNode()
	if err != nil {
		logrus.Warnf("RegisterNode Failed [err:%s]",err.Error())
	}
}
func CronInit() {
	services.InitCrontab()
	services.RunCrontab()

	task.InitSysHealthCheck()

	task.InitLogProcess()
}
func main() {

	// 启动任务框架
	go CronInit()

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
