package task

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ZCache/global"
	"ZCache/services"
	"ZCache/types"
)

// 系统健康检查
func InitSysHealthCheck() {
	spec := global.Config.SysHealthCheckCronSpec
	err := services.AddCrontab(spec, DoSysHealthCheck)
	if err != nil {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoSysHealthCheck"))
	} else {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoSysHealthCheck"))
	}
}


func DoSysHealthCheck(){
	nodes , err := services.GetWorkingNode()
	if err != nil {
		logrus.Warnf("services.GetWorkingNode Failed[err:%s]",err.Error())
	}

	if global.Config.TotalNodes <= nodes{
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_HEALTH
	} else if global.Config.TotalNodes > nodes && global.Config.TotalNodes/2 +1 < nodes{
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_SUBHEALTH
	} else {
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_UNHEALTH
	}

}