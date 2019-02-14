package routes

import (
	"zCache/data"
	"zCache/services"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
)

func DeleteAll() (err error) {
	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_DELETE)
	if err != nil || auth != true {
		return
	}

	// 先存库后删除
	logrus.Infof("%s  export\n", tool.GetFileNameLine())
	err = zdata.CoreFlush()
	if err != nil {
		return
	}

	err = zdata.CoreDeleteAll()

	return
}
