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

func Flush(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	logrus.Infof("%s  Flush\n", tool.GetFileNameLine())
	err = zdata.CoreFlush()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}

}
