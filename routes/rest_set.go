package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestSet(context *gin.Context) {

	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s Set Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)

	err := Set(key, value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success"})
	}
}
