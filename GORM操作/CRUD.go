package main

import "time"

func CreateTable() {
	t := Teacher{
		Name:     "nick",
		Age:      40,
		Roles:    []string{"普通用户", "讲师"},
		birthday: time.Now().Unix(),
		Salary:   12345.1234,
		Email:    "123@qq.com",
	}

	db.Create(t)
}