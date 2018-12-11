package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/global"
	"ZCache/data"
	"ZCache/services"
)

func Delete(context *gin.Context){
	key := context.Param("key")
	logrus.Infof("Delete key: %s",key)
	if nil == global.GlobalVar.Root{
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	}

	//TODO  生成hashIndex
	_ , err := services.GetHashIndex(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"value":"","status":"done"})
		return
	}

	node , err := zdata.Delete(global.GlobalVar.Root, key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		global.GlobalVar.Root = node
		context.JSON(http.StatusOK,gin.H{"key":key, "status":"done"})
		return
	}
}
