package routes

import (
	"zCache/data"
	"zCache/services"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
	"strconv"
)

func Expension(size string) (err error) {
	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		return
	}

	isize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return
	}
	err = zdata.CoreExpension(isize)
	return
}
