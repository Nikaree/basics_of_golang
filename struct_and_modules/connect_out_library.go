package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Программа запущена")
	logrus.Warn("Внимание")
	logrus.Error("Произошла ошибка")
}
