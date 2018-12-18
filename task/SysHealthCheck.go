package task

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ZCache/global"
	"ZCache/services"
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

}