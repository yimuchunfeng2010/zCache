package internal

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
)

func Decr(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s Decr Key:%s\n", tool.GetFileNameLine(), key)

	// TODO
	//if err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "reason": err.Error()})
	//} else {
	//	context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
	//}
}
