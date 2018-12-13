package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool/logrus"
	"ZCache/tool"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()

	logrus.Infof("%s Get All\n", tool.GetFileNameLine())
	node, err := zdata.CoreGetAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"value": "", "status": "done"})
		return
	} else {
		// 遍历链表获取全部数据
		data := make([]types.KeyValue, 0)
		curNode := node

		for nil != curNode {
			data = append(data, types.KeyValue{Key: curNode.Key, Value: curNode.Value})
			curNode = curNode.Next
		}
		context.JSON(http.StatusOK, gin.H{"data": data, "status": "done"})
		return
	}

}
