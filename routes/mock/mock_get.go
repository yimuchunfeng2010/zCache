package mock

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool/logrus"
	"ZCache/tool"
)

func Get(key string) (value string, err error) {
	lockName, err := services.RLock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.RUnlock(lockName)
	logrus.Infof("%s  Get key: %s\n", tool.GetFileNameLine(), key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		return "", err
	}
	return node.Value, nil

}
