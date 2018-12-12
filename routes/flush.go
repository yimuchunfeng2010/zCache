package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Flush(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	logrus.Infof("%s  Flush\n", services.GetFileNameLine())
	err := zdata.CoreFlush()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}

}
