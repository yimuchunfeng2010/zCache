package routes

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Get(context *gin.Context) {
	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	key := context.Param("key")
	logrus.Infof("Get key: %s", key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"value": "", "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}

}
