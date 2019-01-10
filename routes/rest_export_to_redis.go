package routes

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestExportToRedis(context *gin.Context) {
	logrus.Infof("%s  ExportToRedis: %s\n", tool.GetFileNameLine())

	err := ExportToRedis()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "done"})
	}

	return
}
