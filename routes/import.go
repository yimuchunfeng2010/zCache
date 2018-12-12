package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool/logrus"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Import(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	err := zdata.CoreImport()
	if err != nil {
		logrus.Warningf(fmt.Sprintf("Set Failed! [Err:%s]", err.Error()))
		context.JSON(http.StatusOK, gin.H{ "status": "fail"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}
}
