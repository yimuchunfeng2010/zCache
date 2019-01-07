package cluster_inter

import (
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")

	logrus.Infof("%s Update Key:%s\n", tool.GetFileNameLine(), key)
	commitID, err := tool.GetHashIndex("Update" + key + value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Status", "Data": err.Error()})
		return
	}
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_PUT,
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
