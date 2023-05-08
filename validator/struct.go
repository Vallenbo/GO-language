package main

import "github.com/go-playground/validator/v10"

type User struct {
	Name        string            `v:"required,alphaunicode"`
	Age         uint8             `v:"gte=10,lte=130"`
	Phone       string            `v:"required,e164"`
	Email       string            `v:"required,email"`
	Address     *Address          `v:"required"`
	ContactUser []*ContactUser    `v:"required,gte=1,dive"`
	Hobby       []string          `v:"required,gte=2,dive,required,alphaunicode"`
	Data        map[string]string `v:"required,gte=2,dive,keys,len=2,alpha,endkeys,required,alphaunicode"`
}
type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=10,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Email Phone"`
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

func StructVaLidation() {
	v := valiate
	address := &Address{
		Province: "湖南",
		City:     "长沙",
	}
	contactUser1 := &ContactUser{
		Name:    "张三",
		Age:     18,
		Phone:   "+8613800138000",
		Address: address,
	}
	contactUser2 := &ContactUser{
		Name: "李四",
	}
	user := &User{
		Name:        "nick",
		Age:         18,
		Phone:       "+8613800138000",
		Email:       "nick@0voice.com",
		Address:     address,
		ContactUser: []*ContactUser{contactUser1, contactUser2},
		Hobby:       []string{"羽毛球", "乒乓球"},
		Data: map[string]string{
			"AB": "羽毛球",
			"CD": "乒乓球",
		},
	}

	err := v.Struct(user)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err = range errors {
				print(err)
			}
		}
	}
}