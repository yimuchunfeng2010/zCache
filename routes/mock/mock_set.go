package mock

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool/logrus"
	"ZCache/tool"
)

func Set(key string, value string) error {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	logrus.Infof("%s  Set Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)

	_, err := zdata.CoreAdd(key, value)
	if err != nil {
		return err
	}

	return nil
}
