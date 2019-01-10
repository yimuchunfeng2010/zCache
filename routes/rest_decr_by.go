package routes

import (
	"ZCache/client"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func DecrBy(context *gin.Context) {
	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_POST)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		context.JSON(http.StatusOK, gin.H{"status": "done", "reason": err.Error()})
		return

	}
	defer services.Unlock(lockName)
	key := context.Param("key")
	value := context.Param("value")
	step, err := tool.GetContraryNumber(value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "reason": err.Error()})
		return
	}

	logrus.Infof("%s DecrBy Key:%s, step\n", tool.GetFileNameLine(), key, step)
	// 发起提议
	commitID, err := tool.GetHashIndex("DecrBy" + key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
		return
	}
	ackChan := make(chan int64)
	for _, ipAddrPort := range global.Config.ClusterServers {
		go client.GetDecrByAck(ipAddrPort, key, value, ackChan)
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

	// 提交
	if ackCount == len(global.Config.ClusterServers) {
		for _, ipAddrPort := range global.Config.ClusterServers {
			go client.CommitJob(ipAddrPort, commitID)
		}
	} else { //撤销任务
		for _, ipAddrPort := range global.Config.ClusterServers {
			go client.DropJob(ipAddrPort, commitID)
		}
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}
}
