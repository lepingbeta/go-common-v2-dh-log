package dhlog

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("测试info日志")
}

func TestError(t *testing.T) {
	Error("测试error日志")
}
