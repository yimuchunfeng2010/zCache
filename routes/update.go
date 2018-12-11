package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	Data "ZCache/data"
	"ZCache/global"
	"net/http"
)

func Update(context *gin.Context){
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Update Key:%s, Value:%s",key,value)

	data := Data.CacheData{Key:key,Value:value}
	node , err := Data.Update(global.GlobalVar.Root, key, data)
	if err != nil {
		context.JSON(http.StatusConflict,gin.H{"key":key,"value":value, "status":"done"})
		return
	} else {
		global.GlobalVar.Root = node
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}
}

