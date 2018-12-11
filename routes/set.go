package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"

)

func Set(context *gin.Context){
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s",key,value)

	//TODO  生成hashIndex
	_ , err := services.GetHashIndex(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"value":"","status":"done"})
		return
	}

	data := zdata.CacheData{Key:key,Value:value}
	node , err := zdata.Add(global.GlobalVar.Root, key, data)
	if err != nil {
		context.JSON(http.StatusConflict,gin.H{"key":key,"value":value, "status":"done"})
		return
	} else {
		global.GlobalVar.Root = node
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}
}
