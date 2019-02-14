package routes

import (
	"zCache/data"
	"zCache/services"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
)

func Get(key string) (value string, err error) {
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
	logrus.Infof("%s  Get key: %s\n", tool.GetFileNameLine(), key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		return "", err
	} else {
		return node.Value, nil
	}

	return
}
