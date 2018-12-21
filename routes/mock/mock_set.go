package mock

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool/logrus"
	"ZCache/tool"
)

func Set(key string, value string) error {
	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return err

	}
	defer services.RUnlock(lockName)
	logrus.Infof("%s  Set Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)

	_, err = zdata.CoreAdd(key, value)
	if err != nil {
		return err
	}

	return nil
}
