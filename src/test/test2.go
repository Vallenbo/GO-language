package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	dirPath := "F:\\1"                    //指定目录
	fileNames, err := os.ReadDir(dirPath) // 获取文件夹下所有文件名
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`\](\d+)`) // 定义正则表达式
	for i, v := range fileNames {       // 遍历所有文件

		if !fileNames[i].IsDir() { // 如果是文件，而不是文件夹
			content, err := os.ReadFile(dirPath + "/" + v.Name()) // 读取文件内容
			if err != nil {
				fmt.Println(err)
				continue
			}
			newContent := re.ReplaceAll(content, []byte("]")) // 使用正则表达式替换内容

			err = os.WriteFile(dirPath+"/"+v.Name(), newContent, 0644) // 写入修改后的内容
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	fmt.Println("Done!")
}
