package cluster_inter

import (
"ZCache/tool"
"ZCache/global"
"ZCache/tool/logrus"
"github.com/gin-gonic/gin"
"strconv"
	"net/http"
)

func Drop(context *gin.Context) {
	commitID, _ := strconv.ParseInt(context.Param("commitID"), 10, 64)
	logrus.Infof("%s drop ID:%s\n", tool.GetFileNameLine(), commitID)

	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()
	curNode := global.GlobalVar.GPreDoReqList
	preNode := global.GlobalVar.GPreDoReqList
	if curNode != nil {
		if curNode.CommitID == commitID {
			global.GlobalVar.GPreDoReqList.Next = curNode.Next
		} else {
			curNode = curNode.Next
		}
	}

	for curNode != nil {
		if curNode.CommitID == commitID {
			preNode.Next = curNode.Next
			break
		}
		preNode = curNode
		curNode = curNode.Next
	}
	if curNode == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": "Job Not Founc"})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success","Data":commitID})
	}

}