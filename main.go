package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetExe() (string, error) {
	ex, err := os.Executable()
	return ex, err
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	ex, err := GetExe()

	if err == nil {
		logrus.Infof("%v", ex)
	} else {
		logrus.Errorf("%v", err)
	}
}
