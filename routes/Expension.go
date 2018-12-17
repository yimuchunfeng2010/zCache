package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Expension(context *gin.Context) {
	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	size := context.Param("size")
	logrus.Infof("%s  expension size %s\n", tool.GetFileNameLine(), size)

	isize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "fail", "reason": err.Error()})
		return
	}
	err = zdata.CoreExpension(isize)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "fail", "reason": err.Error()})
	}

}
