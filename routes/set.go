package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/types"
	"ZCache/global"

	"ZCache/data"
)

func Set(context *gin.Context){
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s",key,value)

	data := types.CacheData{Key:key,Value:value}
	node , err := zdata.CoreAdd(key, data)
	if err != nil {
		context.JSON(http.StatusConflict,gin.H{"key":key,"value":value, "status":"done"})
		return
	} else {
		global.GlobalVar.Root = node
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}
}
