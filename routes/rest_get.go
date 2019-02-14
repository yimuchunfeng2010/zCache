package routes

import (
	"zCache/data"
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestGet(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s  Get key: %s\n", tool.GetFileNameLine(), key)

	value, err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail"})
	} else {
		context.JSON(http.StatusOK, gin.H{"value": value, "status": "done"})
	}

}
