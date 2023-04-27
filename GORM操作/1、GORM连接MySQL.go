package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB
var err error

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func initDB(dsn string) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info), //打印sql语句
		PrepareStmt: true,                                //预编译sql语句，不支持嵌套事务
	},
	)
	CheckPrintErr(err)
	setPool(db)
}

func CheckPrintErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	CheckPrintErr(err)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(10)
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //postgres数据库
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}) //sqlite 数据库
	//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) //sqlserver 数据库

	dsn := "root:123456@tcp(192.168.4.5:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection mysql err :", err)
	}

	db.AutoMigrate(&Product{}) // 迁移(创建) 表Product

	db.Create(&Product{Code: "D42", Price: 100}) // Create添加数据

	var product Product                   // Read
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Delete(&product, 1) // Delete - 删除 product

}