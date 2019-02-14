package cluster_inter

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Decr(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s Decr Key:%s\n", tool.GetFileNameLine(), key)

	commitID, err := tool.GetHashIndex("Decr" + key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		return
	}
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_DECR,
		Key:      key,
		Value:    "",
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": commitID})
	}
}
