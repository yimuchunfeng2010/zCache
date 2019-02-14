package task

import (
	"zCache/global"
	"zCache/services"
	"fmt"
	"github.com/sirupsen/logrus"
	"zCache/types"
	"time"
)

// 清理过期commit
func CleanOverdueCommit() {
	spec := global.Config.CleanOverdueCommitCronSpec
	err := services.AddCrontab(spec, DoCleanOverdueCommit)
	if err != nil {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoCleanOverdueCommit"))
	} else {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoCleanOverdueCommit"))
	}
}

func DoCleanOverdueCommit() {
	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()
	curNode := global.GlobalVar.GPreDoReqList
	var preNode *types.ProcessingRequest = nil
	for curNode != nil {
		if time.Now().Sub(curNode.CreateTime) > 10 * time.Millisecond{
			if preNode == nil{
				preNode = curNode.Next
				curNode = curNode.Next
			}else{
				preNode.Next = curNode.Next
				curNode = curNode.Next
			}
		} else{
			preNode = curNode
			curNode = curNode.Next
		}

	}

	return
}
