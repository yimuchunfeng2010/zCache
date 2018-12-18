package routes
import (
	"ZCache/global"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/external_data"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImportFromRedis(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()

	err = external_data.ImportFromRedis()
	if err != nil {
		logrus.Warningf("%s ImportFromRedis Failed! [Err:%s]", tool.GetFileNameLine(),err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done","reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "done"})
	}
}
