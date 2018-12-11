package mock

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/sirupsen/logrus"
)

func Set(key string, value string) error {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	logrus.Infof("Set Key:%s, Value:%s", key, value)

	_, err := zdata.CoreAdd(key, value)
	if err != nil {
		return err
	}

	return nil
}
