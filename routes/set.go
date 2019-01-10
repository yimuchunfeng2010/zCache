package routes

import (
	"ZCache/client"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"time"
)

func Set(key string, value string) (err error) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_SET)
	if err != nil || auth != true {
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)
	logrus.Infof("%s Set Key:%s, Value:%s\n", tool.GetFileNameLine(), key, value)

	// 发起提议
	commitID, err := tool.GetHashIndex("Set" + key + value)
	if err != nil {
		return
	}
	ackChan := make(chan int64)
	for _, ipAddrPort := range global.Config.ClusterServers {
		go client.GetSetAck(ipAddrPort, key, value, ackChan)
	}

	timeout := global.Config.Timeout
	ackCount := 0
	for timeout != 0 && ackCount < len(global.Config.ClusterServers) {

		select {
		case _, ok := <-ackChan:
			if ok {
				ackCount++
			}
		default:

		}

		time.Sleep(time.Second / 1000)
		timeout--
	}
	close(ackChan)

	logrus.Infof("commitID %+v", commitID)
	// 提交
	if ackCount == len(global.Config.ClusterServers) {
		for _, ipAddrPort := range global.Config.ClusterServers {
			// todo 获取执行结果
			go client.CommitJob(ipAddrPort, commitID)
		}
	} else { //撤销任务
		for _, ipAddrPort := range global.Config.ClusterServers {
			go client.DropJob(ipAddrPort, commitID)
		}
	}
	return
}
