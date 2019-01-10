package routes

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestDelete(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s  Delete key: %s\n", tool.GetFileNameLine(), key)

	err := Delete(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}
}
