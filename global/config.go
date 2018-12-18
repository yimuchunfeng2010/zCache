package global

import (
	"ZCache/types"
	"sync"
	"os"
)

var Config = struct {
	MaxLen       int64
	UserEmail    string
	MailPort     string
	MailAuthCode string
	MailSmtpHost string
	ToMail       string

	SysHealthCheckCronSpec string
	LogProcessCronSpec string
}{
	MaxLen:       1024,
	UserEmail:    "123456789@qq.com",
	MailPort:     "587",
	MailAuthCode: "mnkxcahklkrebfbb",
	MailSmtpHost: "smtp.qq.com",
	ToMail:       "987654321@qq.com",

	SysHealthCheckCronSpec:"0 */10 * * * *",
	LogProcessCronSpec:"0 */10 * * * *",
}

var GlobalVar = struct {
	GClusterHealthState types.ClusterHealthType
	Root    *types.Node
	GRoot   []*types.Node
	GRootTmp   []*types.Node
	GRWLock *sync.RWMutex
	GCoreInfo types.CoreInfo
	SigChan chan os.Signal
	// 日志指针
	GLogInfoHead *types.LogInfoNode
	GLogInfoTail *types.LogInfoNode
	GLogWarningHead *types.LogInfoNode
	GLogWarningTail *types.LogInfoNode
	GLogErrorHead *types.LogInfoNode
	GLogErrorTail *types.LogInfoNode
}{
	GClusterHealthState:types.CLUSTER_HEALTH_TYPE_HEALTH,
	Root:    nil,
	GRoot:   nil,
	GRootTmp:   nil,
	GRWLock: nil,
	SigChan:nil,
	GCoreInfo:types.CoreInfo{
		KeyNum:0,
	},
	GLogInfoHead:nil,
	GLogInfoTail:nil,
	GLogWarningHead:nil,
	GLogWarningTail:nil,
	GLogErrorHead:nil,
	GLogErrorTail:nil,
}
