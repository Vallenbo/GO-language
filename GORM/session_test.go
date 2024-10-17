package GORM

import (
	"gorm.io/gorm"
)

func Session() {
	tx := db.Session(&gorm.Session{
		PrepareStmt:              true,
		SkipHooks:                true,
		DisableNestedTransaction: true,
		//Logger:                   db.Logger.LogMode(logger.Info),
	})
	tx.Create(&t) //插入数据
}
