package main

import (
	"fmt"

	"os"
	"path/filepath"
	"regexp"
)

func main() {
	dir := "F:\\1" // 指定文件夹路径
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("读取文件夹失败：", err)
		return
	}
	//pattern := regexp.MustCompile(`]`) // 定义正则表达式
	pattern := regexp.MustCompile(`】[0-9]+`) // 定义正则表达式
	for i, _ := range files {
		fmt.Printf("index:%v value.Name:%v\n", i, files[i].Name()) //打印目录
		if files[i].IsDir() {                                      // 跳过文件夹
			continue
		}
		oldName := files[i].Name()
		newName := pattern.ReplaceAllString(oldName, "】") // 删除数字
		fmt.Println(newName)
		if newName != oldName {
			oldPath := filepath.Join(dir, oldName)
			newPath := filepath.Join(dir, newName)
			err = os.Rename(oldPath, newPath) // 重命名文件
			if err != nil {
				fmt.Printf("重命名文件 %s 失败：%v\n", oldName, err)
			} else {
				fmt.Printf("重命名文件 %s 为 %s 成功！\n", oldName, newName)
			}
		}
	}
}
