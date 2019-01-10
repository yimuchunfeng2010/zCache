package routes

import (
	"ZCache/data"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RestExpension(context *gin.Context) {
	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	size := context.Param("size")
	logrus.Infof("%s  expension size %s\n", tool.GetFileNameLine(), size)

	isize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "fail", "reason": err.Error()})
		return
	}
	err = zdata.CoreExpension(isize)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "fail", "reason": err.Error()})
	}

}
