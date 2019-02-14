package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestGetAll(context *gin.Context) {

	logrus.Infof("%s Get All\n", tool.GetFileNameLine())
	data, err := GetAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Value": "", "Status": "done"})
	} else {
		context.JSON(http.StatusOK, gin.H{"Data": data, "Status": "done"})
	}
	return
}
