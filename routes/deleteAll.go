package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteAll(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

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
