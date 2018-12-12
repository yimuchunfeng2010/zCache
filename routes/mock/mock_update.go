package mock

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
)

func Update(key string, value string) error {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	logrus.Infof("%s  Update Key:%s, Value:%s\n", services.GetFileNameLine(), key, value)
	_, err := zdata.CoreUpdate(key, value)
	if err != nil {
		return err
	}

	return nil
}
