package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// 文件内去除字符串.go
func Test_fileDelStr(*testing.T) {
	dir := "E:\\极客时间【专栏】\\007-赵成的运维体系管理课" // 替换为您要处理的目录路径
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("访问文件时出错：%q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" {
			err = replaceInFile(path)
			if err != nil {
				log.Printf("在文件中进行替换时出错：%q: %v\n", path, err)
				return err
			}
			fmt.Printf("已处理文件：%q\n", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func replaceInFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	reg := regexp.MustCompile(`images/.*?/`)
	modifiedData := reg.ReplaceAll(data, []byte("images/"))

	err = os.WriteFile(filename, modifiedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
