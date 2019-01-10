package routes

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestUpdate(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s Update Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)
	err := Update(key, value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}
}
