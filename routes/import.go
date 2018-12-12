package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Import(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	err := zdata.CoreImport()
	if err != nil {
		logrus.Warningf("%s Import Data Failed! [Err:%s]\n", services.GetFileNameLine(), err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "fail"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}
}
