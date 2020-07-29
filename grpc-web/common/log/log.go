package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	file, err := openFile("grpc-web/log/golang-example-app.log")
	if err != nil {
		panic(err)
	}

	logrus.SetOutput(io.MultiWriter(file))
	logrus.SetLevel(logrus.InfoLevel)
}

func openFile(path string) (*os.File, error) {
	_, errStat := os.Stat(path)
	if !os.IsNotExist(errStat) {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0755)
		return f, err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	return f, err
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}
