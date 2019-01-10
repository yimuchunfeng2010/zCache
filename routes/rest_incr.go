package routes

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestIncr(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s Incr:%s,\n", tool.GetFileNameLine(), key)

	err := Incr(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}

	return
}
