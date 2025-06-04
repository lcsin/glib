package ilogrus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	hook, err := NewErrorHook()
	if err != nil {
		panic(err)
	}

	logrus.AddHook(hook)
	logrus.Error("err log ...")
}

func TestCaller(t *testing.T) {
}
