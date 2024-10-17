package GORM

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Roles []string
type Teacher struct {
	Name     string  `gorm:"type:varchar(20);not null"`
	Email    string  `gorm:"type:varchar(20);not null"`
	Salary   float64 `gorm:"type:float"` // 设置为浮点数
	Age      uint8   `gorm:"type:int,check:age>30"`
	birthday int64   `gorm:"serializer:unixtime;type:time"`
	Roles    Roles   `gorm:"type:json"` // 修改为 json 类型
}

func (t *Roles) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Roles) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (Teacher) TableName() string {
	return "teacher"
}

type Product struct {
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "product"
}
