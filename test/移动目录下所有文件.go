package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 指定目录路径
	dir := "E:\\极客时间【专栏】\\018-持续交付36讲\\images"

	// 存储目录名的切片
	var directories []string

	// 获取指定目录下的所有子目录名
	err := getSubDirectories(dir, &directories)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历每个目录名
	for _, directory := range directories {
		// 获取目录下的所有文件
		files, err := os.ReadDir(directory)
		if err != nil {
			log.Fatal(err)
		}

		// 遍历每个文件
		for _, file := range files {
			// 构建源文件路径和目标文件路径
			srcPath := filepath.Join(directory, file.Name())
			destPath := filepath.Join(dir, file.Name())

			// 移动文件
			err := os.Rename(srcPath, destPath)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Moved file: %s\n", srcPath)
		}
	}
}

// 获取指定目录下的所有子目录名
func getSubDirectories(dir string, directories *[]string) error {
	// 获取目录下的所有文件和目录
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// 遍历每个文件和目录
	for _, entry := range entries {
		// 如果是目录，则将目录名添加到切片中
		if entry.IsDir() {
			*directories = append(*directories, filepath.Join(dir, entry.Name()))
		}
	}

	return nil
}
