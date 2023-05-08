package main

import (
	"fmt"
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
var s = &Teacher{Name: "李白", Age: 31}

func Select() {
	rest := db.First(&Teacher{})
	CheckPrintErr(rest.Error)
	fmt.Println("Table rest:", rest)

	db.Where("age>?", 30).First(t).First(&Teacher{}) //将筛选的第一个结果赋值给t
}

func Update() {
	rest := db.Save(&s)
	CheckPrintErr(rest.Error)
	fmt.Println("Table rest:", rest)
}
func CreateTable() {
	err = db.Migrator().AutoMigrate(&Teacher{})
	CheckPrintErr(err)
}

func InsertTable() {
	rest := db.Create(&t) //向表中插入值，在值的 id 中返回插入数据的主键
	CheckPrintErr(rest.Error)
	fmt.Println("Table rest:", rest.RowsAffected, "err:", rest)
}

func Delete() {

	db.Delete(&s)
}