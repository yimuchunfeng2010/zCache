package cluster_inter

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
)

func Commit(context *gin.Context) {
	commitID := context.Param("commitID")
	logrus.Infof("%s commitID Key:%s\n", tool.GetFileNameLine(), commitID)

	//if err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{"status": "NoACK", "reason": err.Error()})
	//} else {
	//	context.JSON(http.StatusOK, gin.H{"status": "ACK","commitID":commitID})
	//}
}
