package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestImportFromRedis(context *gin.Context) {
	err := ImportFromRedis()
	if err != nil {
		logrus.Warningf("%s ImportFromRedis Failed! [Err:%s]", tool.GetFileNameLine(), err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "done"})
	}
}
