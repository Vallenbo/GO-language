package io

import (
	"fmt"
	"os/user"
	"testing"
)

//var err error

func checkErr(err error) {
	if err != nil {
		fmt.Println("To fetch user information is err :", err)
		return
	}
}

func Test_Userinfo(*testing.T) {
	u, err := user.Current()
	checkErr(err)
	fmt.Println("user information :", u.Name, u.Uid, u.HomeDir, u.Gid, u.Username)

	u, err = user.Lookup("LNB\\98560") //查找按用户名查找用户。没有则返回错误
	checkErr(err)
	fmt.Println("user information :", u.Name, u.Uid, u.HomeDir, u.Gid, u.Username)

}
