package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool/logrus"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof(fmt.Sprintf("Delete key: %s", key))
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	_, err := zdata.CoreDelete(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"value": "", "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": key, "status": "done"})
		return
	}
}
