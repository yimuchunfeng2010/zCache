package routes

import (
	"ZCache/data"
	"ZCache/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Flush(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	err := zdata.CoreFlush()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}

}
