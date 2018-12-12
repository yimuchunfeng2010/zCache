package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Set(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s Set Key:%s, Value:%s\n", services.GetFileNameLine(), key, value)

	node, err := zdata.CoreAdd(key, value)
	if err != nil {
		logrus.Warningf("%s Set Failed! [Key:%s, Err:%s]", services.GetFileNameLine(), key, err.Error())
		context.JSON(http.StatusConflict, gin.H{"key": key, "value": value, "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}
}
