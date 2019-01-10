package routes

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
)

func GetKeyNum() (num int, err error) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
		return
	}

	lockName, err := services.RLock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.RUnlock(lockName)

	num, err = zdata.CoreGetKeyNum()
	return

}
