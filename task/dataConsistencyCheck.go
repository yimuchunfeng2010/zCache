package task

import (
	"zCache/data"
	"zCache/global"
	"zCache/services"
	"zCache/tool"
	"fmt"
	"github.com/sirupsen/logrus"
)

// 系统健康检查
func DataConsitencyCheck() {
	spec := global.Config.DataConsitencyCheckCronSpec
	err := services.AddCrontab(spec, DoDataConsitencyCheck)
	if err != nil {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoDataConsitencyCheck"))
	} else {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoDataConsitencyCheck"))
	}
}

func DoDataConsitencyCheck() {
	lockName, err := services.RLock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer services.RUnlock(lockName)

	node, err := zdata.CoreGetAll()
	if err != nil {
		return
	} else {
		// 遍历链表获取全部数据
		curNode := node
		msg := ""
		for nil != curNode {
			msg += curNode.Key+ curNode.Value
			curNode = curNode.Next
		}

		tmp ,err :=tool.GetHashIndex(msg)
		if err != nil {
			logrus.Warningf("GetHashIndex Failed! [Err:%s]", err.Error())
			return
		}
		global.GlobalVar.DataHashIndex = tmp
	}

	return
}
