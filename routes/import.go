package routes

import (
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Import(context *gin.Context) {
	auth, err := services.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	err = zdata.CoreImport()
	if err != nil {
		logrus.Warningf("%s Import Data Failed! [Err:%s]\n", services.GetFileNameLine(), err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "fail"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}
}
