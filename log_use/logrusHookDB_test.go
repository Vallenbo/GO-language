package log_use

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type LogEntry struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement" comment:"自增id" json:"id"`
	Key     string `gorm:"column:key;type:varchar(255)" comment:"日志记录匹配 索引关键值" json:"key"`
	Level   string `gorm:"column:level;type:varchar(255)" comment:"日志level字段" json:"level"`
	Message string `gorm:"column:message;type:text" comment:"日志内容" json:"message"`
}

func (LogEntry) TableName() string {
	return "logToDB"
}

func Test_logrusHookDB(*testing.T) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.5.5:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("无法连接数据库：%v", err)
	}

	// 创建日志钩子，将日志同时写入数据库
	logrus.AddHook(&DatabaseHook{db: db})

	logrus.WithField("key", "example_key").Info("这是一条信息日志")
	logrus.WithField("key", "another_key").Warn("这是一条警告日志")
}

type DatabaseHook struct {
	db *gorm.DB
}

func (d *DatabaseHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (d *DatabaseHook) Fire(entry *logrus.Entry) error {
	fmt.Printf("%+v\n", entry.Data["key"])
	logEntry := LogEntry{
		Key:     entry.Data["key"].(string),
		Level:   entry.Level.String(),
		Message: entry.Message,
	}
	if !d.db.Migrator().HasTable("logToDB") {
		err := d.db.AutoMigrate(&logEntry)
		if err != nil {
			return err
		}
	}

	err := d.db.Create(&logEntry).Error
	if err != nil {
		return err
	}
	return nil
}
