package routes

import (
	"ZCache/external_data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImportFromRedis(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)

	err = external_data.ImportFromRedis()
	if err != nil {
		logrus.Warningf("%s ImportFromRedis Failed! [Err:%s]", tool.GetFileNameLine(), err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "done"})
	}
}
