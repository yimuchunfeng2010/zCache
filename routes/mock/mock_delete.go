package mock

import (
	"zCache/data"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/services"
)

func Delete(key string) error {
	logrus.Infof("%s  Delete key: %s\n", tool.GetFileNameLine(), key)
	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return err

	}
	defer services.Unlock(lockName)
	_, err = zdata.CoreDelete(key)
	if err != nil {
		return err
	}

	return nil
}
