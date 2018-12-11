package mock

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/sirupsen/logrus"
)

func Delete(key string) error {
	logrus.Infof("Delete key: %s", key)
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	_, err := zdata.CoreDelete(key)
	if err != nil {
		return err
	}

	return nil
}
