package routes

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Set(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s", key, value)

	node, err := zdata.CoreAdd(key, value)
	if err != nil {
		logrus.Warningf("Set Failed! [Key:%s, Err:%s]", key, err.Error())
		context.JSON(http.StatusConflict, gin.H{"key": key, "value": value, "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}
}
