package test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// 移动子目录文件至指定目录.go
func Test_mvSubDirToOther(t *testing.T) { //将指定目录及子目录所有文件移动到指定目录下
	sourceDirectory := "E:\\迅雷下载\\216-Vue3企业级项目实战课【完结】\\images" // 指定源目录路径
	destDirectory := "E:\\迅雷下载\\216-Vue3企业级项目实战课【完结】\\images"   // 指定目标目录路径

	err := moveFiles(sourceDirectory, destDirectory)
	if err != nil {
		fmt.Println("Error moving files:", err)
		return
	}

	fmt.Println("Files moved successfully.")
}

func moveFiles(sourceDirectory, destDirectory string) error {
	err := filepath.Walk(sourceDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果当前路径是文件，则进行移动操作
		if !info.IsDir() {
			destPath := filepath.Join(destDirectory, info.Name())
			fmt.Printf("Moving file: %s to %s\n", path, destPath)
			err = moveFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func moveFile(sourcePath, destPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	err = sourceFile.Close() // 关闭源文件句柄
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}
