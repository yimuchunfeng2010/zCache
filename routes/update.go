package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
)

func Update(context *gin.Context){
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Update Key:%s, Value:%s",key,value)
	data := types.CacheData{Key:key,Value:value}
	node , err := zdata.CoreUpdate(key, data)
	if err != nil {
		context.JSON(http.StatusConflict,gin.H{"key":key,"value":value, "status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}
}

