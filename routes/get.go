package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"ZCache/global"
	Data "ZCache/data"
	"net/http"
	"github.com/sirupsen/logrus"
)


func Get(context *gin.Context){
	key := context.Param("key")
	logrus.Infof("Get key: %s",key)
	if nil == global.GlobalVar.Root{
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	}

	node , err := Data.Get(global.GlobalVar.Root, key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}


	fmt.Println(key)
}
