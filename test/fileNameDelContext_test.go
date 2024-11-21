package test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// 文件名去除指定内容.go
func Test_fileNameDelContext(*testing.T) {
	dir := "E:\\迅雷下载\\讲解视频" // 指定文件夹路径
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		pattern := regexp.MustCompile(`【msproject】`) // 匹配正则表达式的内容
		oldName := d.Name()
		newName := pattern.ReplaceAllString(oldName, "") // 将匹配到的内容进行替换
		if newName != oldName {
			oldPath := filepath.Join(path)
			newPath := filepath.Join(filepath.Dir(path), newName)
			err = os.Rename(oldPath, newPath) // 重命名文件
			if err != nil {
				fmt.Printf("重命名文件 %s 失败：%v\n", oldName, err)
			} else {
				fmt.Printf("重命名文件 %s 为 %s 成功！\n", oldName, newName)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败：%v\n", err)
	}
}
