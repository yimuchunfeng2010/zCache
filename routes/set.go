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

func Set(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
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
	logrus.Infof("%s Set Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)

	node, err := zdata.CoreAdd(key, value)
	if err != nil {
		logrus.Warningf("%s Set Failed! [Key:%s, Err:%s]", tool.GetFileNameLine(), key, err.Error())
		context.JSON(http.StatusConflict, gin.H{"key": key, "value": value, "status": "done"})
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
	}
}
