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

func ExportToRedis(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail","reason": err.Error()})
		return
	}

	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done","reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)
	logrus.Infof("%s  ExportToRedis: %s\n", tool.GetFileNameLine())

	err = external_data.ExportToRedis()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{ "status": "done","reason":err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "done"})
	}

}
