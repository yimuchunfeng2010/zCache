package mock

import (
	"ZCache/tool/logrus"
	"ZCache/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//测试文件
func Mock_Set(context *gin.Context) {

	// case 1
	key := "aaa"
	value := "bbb"
	err := Set(key, value)
	if err != nil {
		logrus.Warningf("%s  Set Failed! [Key:%s, Value:%s Err:%s]\n", tool.GetFileNameLine(), key, value, err.Error())
		return
	}

	rsp_value, err := Get(key)
	if err != nil {
		logrus.Warningf("%s   Get Failed! [Key:%s Err:%s]\n", tool.GetFileNameLine(), key, err.Error())
		return
	}

	if 0 != strings.Compare(rsp_value, value) {
		logrus.Warningf("%s   Value Not Equal! [Key:%s Rsp, Value:%sErr:%s]\n", tool.GetFileNameLine(), key, value, rsp_value)
		return
	}
	logrus.Infof("Success to Test Set")
	context.JSON(http.StatusOK, gin.H{"status": "done"})
	return
}
