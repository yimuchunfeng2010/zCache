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

func Decr(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_POST)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"Status": "Fail","Data":""})
		return
	}

	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"Status": "Fail","Data": err.Error()})
		return

	}
	defer services.Unlock(lockName)
	key := context.Param("key")
	logrus.Infof("%s Decr Key:%s\n", tool.GetFileNameLine(), key)
	node, err := zdata.CoreInDecr(key, "-1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "Status": "Success"})
	}
}
