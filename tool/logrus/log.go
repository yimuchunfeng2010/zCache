package logrus

import (
	"fmt"
	"zCache/types"
	"zCache/global"
)

func Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	newNode := new(types.LogInfoNode)
	newNode.Msg = msg
	global.GlobalVar.GLogInfoTail.Next = newNode
	global.GlobalVar.GLogInfoTail = newNode
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	newNode := new(types.LogInfoNode)
	newNode.Msg = msg
	global.GlobalVar.GLogWarningTail.Next = newNode
	global.GlobalVar.GLogWarningTail = newNode
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	newNode := new(types.LogInfoNode)
	newNode.Msg = msg
	global.GlobalVar.GLogErrorTail.Next = newNode
	global.GlobalVar.GLogErrorTail = newNode
}
