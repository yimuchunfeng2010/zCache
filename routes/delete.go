package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/global"
	"ZCache/data"
)

func Delete(context *gin.Context){
	key := context.Param("key")
	logrus.Infof("Delete key: %s",key)
	if nil == global.GlobalVar.Root{
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	}

	_ , err := zdata.CoreDelete(key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":key, "status":"done"})
		return
	}
}
