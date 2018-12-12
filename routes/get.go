package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(context *gin.Context) {
	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()
	key := context.Param("key")
	logrus.Infof("%s  Get key: %s\n", services.GetFileNameLine(), key)

	node, err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"value": "", "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}

}
