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
	"sync"
	"ZCache/routes/cluster_inter"
	"ZCache/zcache_rpc/server"

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

	global.GlobalVar.GInternalLock = new(sync.RWMutex)
}
func CronInit() {
	services.InitCrontab()
	services.RunCrontab()

	task.InitSysHealthCheck()

	task.InitLogProcess()
	task.CleanOverdueCommit()
}
func main() {

	// 启动任务框架
	go CronInit()

	// 启动GRPC
	go server.GrpcInit()
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:key", routes.RestGet)
		v1.DELETE("/:key", routes.RestDelete)
		v1.POST("/:key/:value", routes.RestSet)
		v1.PUT("/:key/:value", routes.RestUpdate)

	}

	v2 := router.Group("/v2")
	{
		v2.GET("/keys", routes.RestGetAll)
		v2.GET("/export", routes.RestFlush)
		v2.PUT("/import", routes.RestImport)
		v2.PUT("/keys", routes.RestDeleteAll)
		v2.PUT("/expension/:size", routes.RestExpension)
	}

	v3 := router.Group("/v3")
	{
		v3.GET("/keys_num",routes.RestGetKeyNum)
		v3.PUT("/incr/:key",routes.RestIncr)
		v3.PUT("/incrBy/:key/:value",routes.RestIncrBy)
		v3.PUT("/decr/:key",routes.RestDecr)
		v3.PUT("/decrBy/:key/:value",routes.RestDecrBy)
	}

	v4 := router.Group("/v4")
	{
		v4.PUT("/import_Redis",routes.RestImportFromRedis)
		v4.GET("/export_Redis",routes.RestExportToRedis)
	}
	test := router.Group("/mock")
	{
		test.PUT("/mock_set", mock.Mock_Set)

	}
	internalPath := router.Group("/internal")
	{
		internalPath.PUT("/commit/:commitID",cluster_inter.Commit)
		internalPath.PUT("/drop/:commitID",cluster_inter.Drop)

		internalPath.POST("/:key/:value",cluster_inter.Set)
		internalPath.DELETE("/:key/",cluster_inter.Delete)
		internalPath.PUT("/incr/:key",cluster_inter.Incr)
		internalPath.PUT("/incrBy/:key/:value",cluster_inter.IncrBy)
		internalPath.PUT("/decr/:key",cluster_inter.Decr)
		internalPath.PUT("/decrBy/:key/:value",cluster_inter.DecrBy)
	}

	router.Run(":8000")

}
