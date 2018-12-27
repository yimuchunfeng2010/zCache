package cluster_inter

import (
	"github.com/gin-gonic/gin"
	"ZCache/tool"
	"net/http"
	"ZCache/types"
)

func DecrBy(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	commitID ,err := tool.GetHashIndex("DecrBy"+key+value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		return
	}
	preReq := types.ProcessingRequest{
		CommitID:commitID,
		Req :types.ReqType_DECRBY,
		Key:key,
		Value:value,
		Next:nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success","Data":commitID})
	}
}