package routes

import (
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"ZCache/tool"
)

func GetKeyNum(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_GET)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.RLock()
	defer global.GlobalVar.GRWLock.RUnlock()

	num, err := zdata.CoreGetKeyNum()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"value": num, "status": "done"})
		return
	}

}
