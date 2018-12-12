package logrus
import("runtime"
"github.com/sirupsen/logrus"
"ZCache/types"
	"time"
	"gopkg.in/gin-gonic/gin.v1/json"
)
func CurrentFile() (string ,int){
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "",-1
	}
	return file, line

}

func PackLogMsg(msg string)(string){
	currentFile, line := CurrentFile()
	fileInfo := currentFile + " " + string(line) + "\n"
	logMsg := types.LogMsg{File:fileInfo,Time:time.Now(),Msg:msg}
	ret, _:= json.Marshal(logMsg)
	return string(ret)
}
// 封装日志接口
func Warningf(msg string){
	logMsg := PackLogMsg(msg)
	logrus.Warningf(logMsg)
}

func Infof(msg string){
	logMsg := PackLogMsg(msg)
	logrus.Infof(logMsg)
}

func Errorf(msg string){
	logMsg := PackLogMsg(msg)
	logrus.Errorf(logMsg)
}