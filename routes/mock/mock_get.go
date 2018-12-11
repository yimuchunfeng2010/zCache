package mock

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/sirupsen/logrus"
)

func Get(key string) (value string, err error) {
	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	logrus.Infof("Get key: %s", key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		return "", err
	}
	return node.Value, nil

}
