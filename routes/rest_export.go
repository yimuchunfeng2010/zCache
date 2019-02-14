package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestFlush(context *gin.Context) {
	logrus.Infof("%s  export\n", tool.GetFileNameLine())
	err := Flush()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure"})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}
	return
}
