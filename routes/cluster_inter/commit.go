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
	"errors"
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
		err := DoCommit(toDORequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
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
	default:
		err = errors.New("Wrong Request Type")
	}

	return
}