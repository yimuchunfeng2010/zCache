package mock

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/sirupsen/logrus"
)

func Update(key string, value string) error {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	logrus.Infof("Update Key:%s, Value:%s", key, value)
	_, err := zdata.CoreUpdate(key, value)
	if err != nil {
		return err
	}

	return nil
}
