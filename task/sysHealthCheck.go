package task

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/types"
	"fmt"
	"github.com/sirupsen/logrus"
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

func DoSysHealthCheck() {
	nodes, err := services.GetWorkingNode()
	if err != nil {
		logrus.Warnf("services.GetWorkingNode Failed[err:%s]", err.Error())
	}

	if global.Config.TotalNodes <= nodes {
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_HEALTH
	} else if global.Config.TotalNodes > nodes && global.Config.TotalNodes/2+1 < nodes {
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_SUBHEALTH
	} else {
		global.GlobalVar.GClusterHealthState = types.CLUSTER_HEALTH_TYPE_UNHEALTH
	}

	// 若集群处于亚健康和不健康状态，则备份数据
	if types.CLUSTER_HEALTH_TYPE_SUBHEALTH == global.GlobalVar.GClusterHealthState || types.CLUSTER_HEALTH_TYPE_UNHEALTH == global.GlobalVar.GClusterHealthState {
		if false == global.GlobalVar.IsAlreadyBackup {
			lockName, err := services.Lock()
			if err != nil {
				logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
				return

			}
			defer services.Unlock(lockName)
			zdata.CoreFlush()
			global.GlobalVar.IsAlreadyBackup = true
		}
	} else if types.CLUSTER_HEALTH_TYPE_HEALTH == global.GlobalVar.GClusterHealthState {
		global.GlobalVar.IsAlreadyBackup = false
	}

	return
}
