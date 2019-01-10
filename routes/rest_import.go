package routes

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestImport(context *gin.Context) {
	err := Import()
	if err != nil {
		logrus.Warningf("%s Import Data Failed! [Err:%s]\n", tool.GetFileNameLine(), err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "fail"})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}
	return
}
