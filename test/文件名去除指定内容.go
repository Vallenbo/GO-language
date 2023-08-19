package main

import (
	"fmt"

	"os"
	"path/filepath"
	"regexp"
)

func main() {
	dir := "E:\\linux性能优化实战" // 指定文件夹路径
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("读取文件夹失败：", err)
		return
	}
	//pattern := regexp.MustCompile(`]`) // 定义正则表达式
	pattern := regexp.MustCompile(`【海量资源：666java.com】`) // 匹配正则表达式的内容 // 匹配：[前不限字符，[]内1-9不限数字
	for i, _ := range files {
		fmt.Printf("the Name of file:%v\n", files[i].Name()) //打印目录
		if files[i].IsDir() {                                // 跳过文件夹
			continue
		}
		oldName := files[i].Name()
		newName := pattern.ReplaceAllString(oldName, "") // 将匹配到的内容进行替换
		//fmt.Println(newName)
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
