package log_use

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type CustomHook struct{}

func (h *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *CustomHook) Fire(entry *logrus.Entry) error {
	entry.Data["custom_field"] = "custom_value" // 在日志中添加额外的字段
	if entry.Level >= logrus.ErrorLevel {       // 根据日志级别执行特定操作
		fmt.Println("这是一个错误日志，执行一些特殊操作...")
	}
	return nil
}

func Test_LogrusHook(t *testing.T) {
	Testlog := logrus.New()

	// 添加自定义 Hook
	hook := &CustomHook{}
	Testlog.AddHook(hook)

	Testlog.WithFields(logrus.Fields{"这是key": "这是value"}).Info("这是一条信息日志")
	Testlog.Error("这是一条错误日志")
}
