package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestIncrBy(context *gin.Context) {

	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s IncrBy Key:%s value %s\n", tool.GetFileNameLine(), key, value)

	err := IncrBy(key, value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}

}
