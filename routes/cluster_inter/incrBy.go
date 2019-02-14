package cluster_inter

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IncrBy(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")

	logrus.Infof("%s IncrBy Key:%s\n", tool.GetFileNameLine(), key)
	commitID, err := tool.GetHashIndex("IncrBy" + key + value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		return
	}
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_INCRBY,
		Key:      key,
		Value:    value,
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": commitID})
	}

}
