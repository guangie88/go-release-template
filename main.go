package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	ex, err := os.Executable()

	if err == nil {
		logrus.Infof("%v", ex)
	} else {
		logrus.Errorf("%v", err)
	}
}
