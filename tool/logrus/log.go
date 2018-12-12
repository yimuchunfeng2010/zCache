package logrus

import (
	"github.com/sirupsen/logrus"
)

func Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
