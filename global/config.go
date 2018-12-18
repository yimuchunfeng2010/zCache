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
}{
	MaxLen:       1024,
	UserEmail:    "123456789@qq.com",
	MailPort:     "587",
	MailAuthCode: "mnkxcahklkrebfbb",
	MailSmtpHost: "smtp.qq.com",
	ToMail:       "987654321@qq.com",

	SysHealthCheckCronSpec:"0 */10 * * * *",
}

var GlobalVar = struct {
	GClusterHealthState types.ClusterHealthType
	Root    *types.Node
	GRoot   []*types.Node
	GRootTmp   []*types.Node
	GRWLock *sync.RWMutex
	GCoreInfo types.CoreInfo
	SigChan chan os.Signal
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
}
