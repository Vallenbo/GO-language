package main

import (
	"fmt"
	"os/user"
)

//var err error

func checkErr(err error) {
	if err != nil {
		fmt.Println("To fetch user information is err :", err)
		return
	}
}

func main() {
	u, err := user.Current()
	checkErr(err)
	fmt.Println("user information :", u.Name, u.Uid, u.HomeDir, u.Gid, u.Username)

	u, err = user.Lookup("LNB\\98560") //查找按用户名查找用户。没有则返回错误
	checkErr(err)
	fmt.Println("user information :", u.Name, u.Uid, u.HomeDir, u.Gid, u.Username)

}
