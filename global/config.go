package global

import (
	"ZCache/types"
	"os"
	"sync"
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
	Port         string

	SysHealthCheckCronSpec      string
	DataConsitencyCheckCronSpec string
	CleanOverdueCommitCronSpec string
	LogProcessCronSpec          string
	ClusterServers              []string
	Timeout                     int
	GrpcPort                    string
}{
	MaxLen:       1024,
	UserEmail:    "123456789@qq.com",
	MailPort:     "587",
	MailAuthCode: "mnkxcahklkrebfbb",
	MailSmtpHost: "smtp.qq.com",
	ToMail:       "987654321@qq.com",

	SysHealthCheckCronSpec:      "0 */10 * * * *",
	LogProcessCronSpec:          "0 */10 * * * *",
	DataConsitencyCheckCronSpec: "0 */10 * * * *",
	CleanOverdueCommitCronSpec:  "0 */10 * * * *",

	TotalNodes:     3,
	ZkIPaddr:       "192.168.228.143:2181",
	Port:           "8000",
	ClusterServers: []string{"127.0.0.1:8000"},
	Timeout:        5000,    //ms
	GrpcPort:       "50051", // grpc端口号
}

var GlobalVar = struct {
	GClusterHealthState types.ClusterHealthType
	Root                *types.Node
	GRoot               []*types.Node
	GRootTmp            []*types.Node
	GCoreInfo           types.CoreInfo
	SigChan             chan os.Signal
	IsAlreadyBackup     bool
	DataHashIndex       int64
	GInternalLock       *sync.RWMutex
	GPreDoReqList       *types.ProcessingRequest
	// 日志指针
	GLogInfoHead    *types.LogInfoNode
	GLogInfoTail    *types.LogInfoNode
	GLogWarningHead *types.LogInfoNode
	GLogWarningTail *types.LogInfoNode
	GLogErrorHead   *types.LogInfoNode
	GLogErrorTail   *types.LogInfoNode
}{
	GClusterHealthState: types.CLUSTER_HEALTH_TYPE_HEALTH,
	Root:                nil,
	GRoot:               nil,
	GRootTmp:            nil,
	SigChan:             nil,
	GCoreInfo: types.CoreInfo{
		KeyNum: 0,
	},
	DataHashIndex:   0,
	IsAlreadyBackup: false,
	GLogInfoHead:    nil,
	GLogInfoTail:    nil,
	GLogWarningHead: nil,
	GLogWarningTail: nil,
	GLogErrorHead:   nil,
	GLogErrorTail:   nil,
	GInternalLock:   nil,
	GPreDoReqList:   nil,
}
