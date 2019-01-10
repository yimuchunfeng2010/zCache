package cluster_inter

import (
	"ZCache/tool"
	"ZCache/global"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"ZCache/types"
	"ZCache/data"
)

func Commit(context *gin.Context) {
	commitID ,_ := strconv.ParseInt(context.Param("commitID"), 10, 64)
	logrus.Infof("%s commit Job ID:%s\n", tool.GetFileNameLine(), commitID)

	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()

	var toDORequest  *types.ProcessingRequest = nil
	if(global.GlobalVar.GPreDoReqList.CommitID == commitID){
		toDORequest = global.GlobalVar.GPreDoReqList
		global.GlobalVar.GPreDoReqList = global.GlobalVar.GPreDoReqList.Next
	} else{
		var tmpPre *types.ProcessingRequest = global.GlobalVar.GPreDoReqList
		for tmpPre.Next != nil && tmpPre.Next.CommitID != commitID{
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && tmpPre.Next.CommitID == commitID{
			toDORequest = tmpPre.Next
			tmpPre.Next = tmpPre.Next.Next
		}

	}


	if toDORequest == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": "commit not found"})
		return
	} else {
		err := zdata.DoCommit(toDORequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
		}else{
			context.JSON(http.StatusOK, gin.H{"Status": "Success","Data":commitID})
		}
	}

	return
}

