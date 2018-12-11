package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ZCache/types"
	"ZCache/data"
	"ZCache/global"
)

func Set(context *gin.Context){
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s",key,value)

	data := types.CacheData{Key:key,Value:value}
	node , err := zdata.CoreAdd(key, data)
	if err != nil {
		logrus.Warningf("Set Failed! [Key:%s, Err:%s]", key,err.Error())
		context.JSON(http.StatusConflict,gin.H{"key":key,"value":value, "status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}
}
