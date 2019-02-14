package routes

import (
	"zCache/global"
	"zCache/services"
	"zCache/tool"
	"zCache/tool/logrus"
	"zCache/types"
	"time"
	Client "zCache/zcache_rpc/client"
	pb "zCache/zcache_rpc/zcacherpc"
)

func Delete(key string) (err error) {
	logrus.Infof("%s  Delete key: %s\n", tool.GetFileNameLine(), key)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_DELETE)
	if err != nil || auth != true {
		return
	}

	lockName, err := services.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.Unlock(lockName)

	logrus.Infof("%s Delete Key:%s\n", tool.GetFileNameLine(), key)
	// 发起提议
	commitID, err := tool.GetHashIndex("Delete" + key)
	if err != nil {
		return
	}
	ackChan := make(chan string)
	for _, client := range  global.Config.Clients{
		go func(){
			Client.PreDelete(client,pb.Data{Key:key},ackChan)
		}()
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
		for _, client := range global.Config.Clients {
			go Client.CommitJob(client, pb.CommitIDMsg{CommitID:string(commitID)})
		}
	} else { //撤销任务
		for _, client := range global.Config.Clients {
			go Client.CommitJob(client, pb.CommitIDMsg{CommitID:string(commitID)})
		}
	}
	return
}
