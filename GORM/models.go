package main

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name     string   `gorm:"size:256"`
	Email    string   `gorm:"size:256"`
	Salary   float64  `gorm:"precision:7"` // 设置为浮点数
	Age      uint8    `gorm:"check:age>30"`
	birthday int64    `gorm:"serializer:unixtime;type:time"`
	Roles    []string `gorm:"type:text;serializer:json"` // 修改为 text 类型
}