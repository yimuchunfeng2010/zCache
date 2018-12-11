package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sirupsen/logrus"
	"ZCache/data"
	"ZCache/global"
)


func Get(context *gin.Context){
	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	key := context.Param("key")
	logrus.Infof("Get key: %s",key)

	node , err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":node.Key,"value":node.Value, "status":"done"})
		return
	}


}
