package mock

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool/logrus"
	"ZCache/tool"
)

func Get(key string) (value string, err error) {
	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	logrus.Infof("%s  Get key: %s\n", tool.GetFileNameLine(), key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		return "", err
	}
	return node.Value, nil

}
