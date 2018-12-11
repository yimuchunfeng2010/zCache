package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ZCache/global"
	"net/http"
	Data "ZCache/data"
)

func Delete(context *gin.Context){
	key := context.Param("key")
	logrus.Infof("Delete key: %s",key)
	if nil == global.GlobalVar.Root{
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	}

	_ , err := Data.Delete(global.GlobalVar.Root, key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":key, "status":"done"})
		return
	}
}
