package routes

import (
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
	"ZCache/tool/logrus"
	"ZCache/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Decr(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_POST)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	logrus.Infof("%s Decr Key:%s\n", tool.GetFileNameLine(), key)
	node, err := zdata.CoreInDecr(key, "-1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail","reason":err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}
}
