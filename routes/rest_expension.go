package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestExpension(context *gin.Context) {

	size := context.Param("size")
	logrus.Infof("%s  expension size %s\n", tool.GetFileNameLine(), size)

	err := Expension(size)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "fail", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}

}
