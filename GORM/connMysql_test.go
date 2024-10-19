package GORM

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
	"time"
)

var dsn = "root:123456@tcp(192.168.5.5:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func Test_createData(T *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //postgres数据库
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}) //sqlite 数据库
	//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) //sqlserver 数据库
	OperatorDB()
	//CreateTable()
	//InsertTable()
}

func OperatorDB() { //官方演示操作
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection mysql err :", err)
	}

	db.AutoMigrate(&Product{}) // 创建 表Product

	db.Create(&Product{Code: "D42", Price: 100}) // Create添加数据

	var product Product   // Read
	db.First(&product, 1) // 根据整型主键查找
	fmt.Printf("产品  :%v\n", product)
	// 查找 code 字段值为 D42 的记录
	err = db.First(&product, "code = ?", "D42").Error
	if err != nil {
		println("查询产品 code err：%v", err)
	}
	println("产品 code :%v\n", product.Code)

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//db.Delete(&product, 1) // Delete - 删除 product
}

func initDB() { //初始化连接
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

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	CheckPrintErr(err)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(10)
}

func CheckPrintErr(err error) {
	if err != nil {
		log.Println("err：", err)
		return
	}
}
