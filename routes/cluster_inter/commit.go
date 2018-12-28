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
	curNode := global.GlobalVar.GPreDoReqList
	preNode := global.GlobalVar.GPreDoReqList
	if curNode != nil{
		if curNode.CommitID == commitID {
			global.GlobalVar.GPreDoReqList.Next = curNode
			goto ProcessCommit
		}else{
			curNode = curNode.Next
		}
	}

	for curNode != nil {
		if curNode.CommitID == commitID{
			preNode.Next = curNode.Next
			goto ProcessCommit
		}
		preNode = curNode
		curNode = curNode.Next
	}

	ProcessCommit:
	if curNode == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": "curNode nil"})
	} else {
		err := DoCommit(curNode)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		}else{
			context.JSON(http.StatusOK, gin.H{"Status": "Success","Data":commitID})
		}
	}

	return
}

func DoCommit(data *types.ProcessingRequest)(err error){

	switch data.Req {
	case types.ReqType_POST:
		_, err = zdata.CoreAdd(data.Key,data.Value)
	case types.ReqType_DELETE:
		_, err = zdata.CoreDelete(data.Key)
	case types.ReqType_PUT:
		_, err = zdata.CoreUpdate(data.Key,data.Value)
	case types.ReqType_INCR,types.ReqType_INCRBY,types.ReqType_DECR,types.ReqType_DECRBY:
		_, err = zdata.CoreInDecr(data.Key,data.Value)
	}
	return
}