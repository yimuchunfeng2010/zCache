package task

import (
	"fmt"
	zLogrus "ZCache/tool/logrus"
	"github.com/sirupsen/logrus"
	"ZCache/global"
	"ZCache/services"
)

// 系统健康检查
func InitLogProcess() {
	spec := global.Config.LogProcessCronSpec
	err := services.AddCrontab(spec, DoLogProcess)
	if err != nil {
		zLogrus.Warningf(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoLogProcess"))
	} else {
		zLogrus.Infof(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoLogProcess"))
	}
}


func DoLogProcess(){

	// Info级别日志
	InfoLogProcess()

	// Warning级别日志
	WarningLogProcess()

	// Error级别日志
	ErrorLogProcess()
}

func InfoLogProcess(){
	curNode := global.GlobalVar.GLogInfoHead.Next
	if nil == curNode{
		zLogrus.Warningf("GLogInfoHead nil")
		return
	}

	for nil != curNode{
		logrus.Infof(curNode.Msg)
		curNode = curNode.Next
	}

	// 更新头节点指针
	global.GlobalVar.GLogInfoHead = curNode
}

func WarningLogProcess(){
	curNode := global.GlobalVar.GLogWarningHead.Next
	if nil == curNode{
		zLogrus.Warningf("GLogWarningHead nil")
		return
	}

	for nil != curNode{
		logrus.Infof(curNode.Msg)
		curNode = curNode.Next
	}

	// 更新头节点指针
	global.GlobalVar.GLogWarningHead = curNode
}

func ErrorLogProcess(){
	curNode := global.GlobalVar.GLogErrorHead.Next
	if nil == curNode{
		zLogrus.Warningf("GLogErrorHead nil")
		return
	}

	for nil != curNode{
		logrus.Infof(curNode.Msg)
		curNode = curNode.Next
	}

	// 更新头节点指针
	global.GlobalVar.GLogErrorHead = curNode
}