package GORM

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

var t = Teacher{
	Name:     "nick",
	Age:      40,
	Roles:    []string{"普通用户", "讲师"},
	birthday: time.Now().Unix(),
	Salary:   12345.1234,
	Email:    "123@qq.com",
}

func Test_crud(*testing.T) {
	//dsn := "root:123456@tcp(192.168.5.5:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(192.168.31.202:3306)/qfwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection mysql err :", err)
	}
	err := db.Table("login").Where("username = test")

	//InsertTable()
	//Select()
	Update()
}

func InsertTable() {
	//err = db.AutoMigrate(&Teacher{})
	if err != nil {
		fmt.Println(" db.AutoMigrate Table err:", err)
		return
	}
	rest := db.Create(&Teacher{
		Name:     "nick",
		Age:      40,
		Roles:    []string{"普通用户", "讲师"},
		birthday: time.Now().Unix(),
		Salary:   12345.1234,
		Email:    "123@qq.com",
	}) //向表中插入值，在值的 id 中返回插入数据的主键
	CheckPrintErr(rest.Error)
	fmt.Println("Table 插入行:", rest.RowsAffected, "err:", rest)
}

var s = &Teacher{Name: "李白", Age: 31}

func Select() {
	//var test []Teacher
	//rest := db.First(&test)

	//var s []string
	//rest := db.Table("teacher").Select("name").Find(&s)

	var results []map[string]interface{}
	rest := db.Table("teacher").Find(&results)

	CheckPrintErr(rest.Error)
	for key, v := range results {
		fmt.Println("key:", key, "value:", v)
	}
	//fmt.Println("db.find row: ", rest.RowsAffected, "db.First  rest:", results)

	//db.Where("age>?", 30).First(t).First(&Teacher{}) //将筛选的第一个结果赋值给t
}

func Update() {
	rest := db.Model(&Teacher{}).Where("id < 3").Updates(Teacher{Age: 30})
	CheckPrintErr(rest.Error)
	fmt.Println("Table rest RowsAffected:", rest.RowsAffected)
}
func CreateTable() {
	err = db.Migrator().AutoMigrate(&Teacher{})
	CheckPrintErr(err)
}

func Delete() {
	db.Delete(&s)
}
