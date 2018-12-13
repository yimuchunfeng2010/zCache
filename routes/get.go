package routes

import (
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"ZCache/tool"
)

func Get(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	key := context.Param("key")
	logrus.Infof("%s  Get key: %s\n", tool.GetFileNameLine(), key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"value": "", "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}

}
