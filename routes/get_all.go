package routes

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
)

func GetAll() (data []types.KeyValue, err error) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)

	logrus.Infof("%s Get All\n", tool.GetFileNameLine())
	node, err := zdata.CoreGetAll()
	if err != nil {
	} else {
		// 遍历链表获取全部数据
		curNode := node

		for nil != curNode {
			data = append(data, types.KeyValue{Key: curNode.Key, Value: curNode.Value})
			curNode = curNode.Next
		}
	}
	return

}
