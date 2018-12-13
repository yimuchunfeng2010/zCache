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
}{
	MaxLen:       1024,
	UserEmail:    "123456789@qq.com",
	MailPort:     "587",
	MailAuthCode: "mnkxcahklkrebfbb",
	MailSmtpHost: "smtp.qq.com",
	ToMail:       "987654321@qq.com",
}

var GlobalVar = struct {
	GClusterHealthState types.ClusterHealthType
	Root    *types.Node
	GRoot   []*types.Node
	GRWLock *sync.RWMutex
	GCoreInfo types.CoreInfo
	SigChan chan os.Signal
}{
	GClusterHealthState:types.CLUSTER_HEALTH_TYPE_HEALTH,
	Root:    nil,
	GRoot:   nil,
	GRWLock: nil,
	SigChan:nil,
	GCoreInfo:types.CoreInfo{
		KeyNum:0,
	},
}
