package mock

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
)

func Delete(key string) error {
	logrus.Infof("%s  Delete key: %s\n", services.GetFileNameLine(), key)
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	_, err := zdata.CoreDelete(key)
	if err != nil {
		return err
	}

	return nil
}
