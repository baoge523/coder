package mysql_driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// https://gorm.io/zh_CN/docs/index.html

// mysql 驱动、获取mysql连接、连接池、mysql的其他信息

// 查看tencent的mysql 监控了哪些数据信息，如果我要在mysql里面如何查看这些监控信息；其他信息也要这样学习一下


func BuildDB() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root123456", "127.0.0.1", "test")
	// root:root123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local
	fmt.Printf("%s \n", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 指定gorm的日志输出级别为info https://gorm.io/zh_CN/docs/logger.html
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("%v \n",err)
		return nil
	}

	return db
}

func mainTest() {
	db := BuildDB()
	var dest []map[string]interface{}
	db.Table("policy").Find(&dest)
	fmt.Printf("%v", dest)
}