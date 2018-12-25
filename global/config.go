package global

import (
	"ZCache/types"
	"os"
)

var Config = struct {
	MaxLen       int64
	UserEmail    string
	MailPort     string
	MailAuthCode string
	MailSmtpHost string
	ToMail       string
	TotalNodes   int
	ZkIPaddr     string

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

	TotalNodes:3,
	ZkIPaddr:"192.168.228.143:2181",
}

var GlobalVar = struct {
	GClusterHealthState types.ClusterHealthType
	Root    *types.Node
	GRoot   []*types.Node
	GRootTmp   []*types.Node
	GCoreInfo types.CoreInfo
	SigChan chan os.Signal
	IsAlreadyBackup bool
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
	SigChan:nil,
	GCoreInfo:types.CoreInfo{
		KeyNum:0,
	},
	IsAlreadyBackup:false,
	GLogInfoHead:nil,
	GLogInfoTail:nil,
	GLogWarningHead:nil,
	GLogWarningTail:nil,
	GLogErrorHead:nil,
	GLogErrorTail:nil,
}
