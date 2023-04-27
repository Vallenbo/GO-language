package main

import "gorm.io/gorm"

type Roles []string
type Teacher struct {
	gorm.Model
	Name     string  `gorm:"size:256"`
	Email    string  `gorm:"size:256"`
	Salary   float64 `gorm:"Salary:256;precision:7"` //设置为浮点数
	Age      uint8   `gorm:"check:age>30"`
	birthday int64   `gorm:"serializer:unixtime;type:time"`
	Roles    Roles   `gorm:"serializer:json"`
}