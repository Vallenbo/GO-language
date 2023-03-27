package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed ,err:", err)
		return
	}
	for line := range tails.Lines {
		fmt.Println("line:", line.Text)
	}
}
