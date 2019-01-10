package routes

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestGetAll(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"Status": "done", "Reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)

	logrus.Infof("%s Get All\n", tool.GetFileNameLine())
	node, err := zdata.CoreGetAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Value": "", "Status": "done"})
	} else {
		// 遍历链表获取全部数据
		data := make([]types.KeyValue, 0)
		curNode := node

		for nil != curNode {
			data = append(data, types.KeyValue{Key: curNode.Key, Value: curNode.Value})
			curNode = curNode.Next
		}
		context.JSON(http.StatusOK, gin.H{"Data": data, "Status": "done"})
	}

}
