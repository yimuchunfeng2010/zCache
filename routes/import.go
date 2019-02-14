package routes

import (
	"zCache/data"
	"zCache/services"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
)

func Import() (err error) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		return
	}
	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)
	err = zdata.CoreImport()
	return
}
