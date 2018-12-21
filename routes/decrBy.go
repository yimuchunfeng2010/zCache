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

func DecrBy(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_POST)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done","reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)
	key := context.Param("key")
	value := context.Param("value")
	step, err := tool.GetContraryNumber(value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "reason": err.Error()})
		return
	}

	logrus.Infof("%s DecrBy Key:%s, step\n", tool.GetFileNameLine(), key, step)
	node, err := zdata.CoreInDecr(key, step)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
	}
}
