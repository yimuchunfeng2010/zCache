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

func DeleteAll(context *gin.Context) {
	lockName, err := services.Lock()
	if err != nil{
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done","reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_DELETE)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	// 先存库后删除
	logrus.Infof("%s  export\n", tool.GetFileNameLine())
	err = zdata.CoreFlush()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure", "reason": err.Error()})
		return
	}

	err = zdata.CoreDeleteAll()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}

}
