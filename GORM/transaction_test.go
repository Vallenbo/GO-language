package GORM

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Animal struct {
	Name string
}

func Test_Transaction(*testing.T) {
	dsn := "root:123456@tcp(192.168.5.5:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection mysql err :", err)
	}

	db.AutoMigrate(&Animal{})
	err = db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err = tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			fmt.Println("db.Transaction tx.Create 1 err : ", err)
			return err
		}
		if err = tx.Delete(&Animal{Name: "sdfd"}).Error; err != nil {
			fmt.Println("db.Transaction tx.Delete 2  err : ", err)
			return err
		}

		// 返回 nil 提交事务
		println("Transaction is ok")
		return nil
	})
	if err != nil {
		fmt.Println("db.Transaction err : ", err)
		return
	}
}
