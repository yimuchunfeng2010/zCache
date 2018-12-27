package internal
import (
	"github.com/gin-gonic/gin"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"net/http"
	"ZCache/types"
)

func Set(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")

	logrus.Infof("%s Set Key:%s\n", tool.GetFileNameLine(), key)
	commitID ,err := tool.GetHashIndex("Set"+key+value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "NoACK", "reason": err.Error()})
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
		context.JSON(http.StatusInternalServerError, gin.H{"status": "NoACK", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "ACK","commitID":commitID})
	}

}