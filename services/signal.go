package services

import (
	"os"
	"syscall"
	"zCache/data"
	"zCache/tool/logrus"
	"zCache/tool"
)

func SigHandler(sigChan chan os.Signal)  {
	select {
	case signal := <- sigChan:{
		switch signal {
		case syscall.SIGHUP,syscall.SIGINT,syscall.SIGQUIT,syscall.SIGKILL:
			logrus.Infof("%s Receive Signal %+v",tool.GetFileNameLine(), signal)
			zdata.CoreFlush()
			return
		default:
			logrus.Warningf("%s Unhanlde Signal %+v",signal)
			return

		}
	}

	}
}
