package mysql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/index.html

// mysql 驱动、获取mysql连接、连接池、mysql的其他信息

// 查看tencent的mysql 监控了哪些数据信息，如果我要在mysql里面如何查看这些监控信息；其他信息也要这样学习一下


func buildDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("mysql"), &gorm.Config{})
	if err != nil {

	}

	return db
}
