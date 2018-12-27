package cluster_inter

import (
	"ZCache/tool"
	"ZCache/global"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Commit(context *gin.Context) {
	commitID ,_ := strconv.ParseInt(context.Param("commitID"), 10, 64)
	logrus.Infof("%s commitID Key:%s\n", tool.GetFileNameLine(), commitID)

	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()
	curNode := global.GlobalVar.GPreDoReqList
	preNode := global.GlobalVar.GPreDoReqList
	if curNode != nil{
		if curNode.CommitID == commitID {
			//TODO
			return
		}else{
			curNode = curNode.Next
		}
	}

	for curNode != nil {
		if curNode.CommitID == commitID{
			preNode.Next = curNode.Next
			// TODO
			return
		}
		preNode = curNode
		curNode = curNode.Next
	}

	//if err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{"status": "NoACK", "reason": err.Error()})
	//} else {
	//	context.JSON(http.StatusOK, gin.H{"status": "ACK","commitID":commitID})
	//}
}
