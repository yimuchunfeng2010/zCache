package services

import (
	"os"
	"syscall"
	"ZCache/data"
	"ZCache/tool/logrus"
	"ZCache/tool"
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
