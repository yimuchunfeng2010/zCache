package routes

import (
	"ZCache/external_data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
)

func ExportToRedis() (err error) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)
	logrus.Infof("%s  ExportToRedis: %s\n", tool.GetFileNameLine())

	err = external_data.ExportToRedis()
	return
}
