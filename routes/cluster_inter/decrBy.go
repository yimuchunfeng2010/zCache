package cluster_inter

import (
	"ZCache/tool"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DecrBy(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	commitID, err := tool.GetHashIndex("DecrBy" + key + value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		return
	}
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_DECRBY,
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
