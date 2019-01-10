package routes

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
)

func Flush() (err error) {
	lockName, err := services.RLock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.RUnlock(lockName)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
	}

	logrus.Infof("%s  export\n", tool.GetFileNameLine())
	err = zdata.CoreFlush()
	return
}
