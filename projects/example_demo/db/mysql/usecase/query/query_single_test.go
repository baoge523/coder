package query

import (
	"fmt"
	mysql_driver "projects/example_demo/db/mysql"
	"testing"
)

/**
GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误

First 按照主键排查、加limit 1

Take 没有排查，加limit 1

Last 按照主键倒排序，加limit 1

 */

// 基于map的方式调用 First、Take、Last : 需要指定table，同时first和map需要排序无法排序，查询数据失败
func TestQuerySingle_map(t *testing.T) {
	db := mysql_driver.BuildDB()
	var dest1 map[string]interface{}
	var dest2 map[string]interface{}
	var dest3 map[string]interface{}

	// sql：SELECT * FROM `policy` ORDER BY `policy`. LIMIT 1
	db.Table("policy").First(&dest1)
	fmt.Printf("first: %v \n", dest1) // map[]

	// sql: SELECT * FROM `policy` LIMIT 1
	db.Table("policy").Take(&dest2)
	fmt.Printf("take: %v \n", dest2) // map[condition_id:4 id:1]

	// sql: SELECT * FROM `policy` ORDER BY `policy`. DESC LIMIT 1
	db.Table("policy").Last(&dest3)
	fmt.Printf("last: %v \n", dest3) // map[]
}

// 模型定义参考: https://gorm.io/zh_CN/docs/models.html
type policy struct {
	Id int64
	ConditionId int64
}

// type Tabler interface 自动实现了该接口
func (p *policy) TableName() string {
	return "policy"
}

// 基于结构体的查询，需要实现 TableName() string 方法，这样可以自动找到对应的表
func TestQuerySingle_struct(t *testing.T) {
	db := mysql_driver.BuildDB()
	var dest policy
	var dest2 policy
	var dest3 policy

	// sql: SELECT * FROM `policy` ORDER BY `policy`.`id` LIMIT 1
	// db.Table("policy").First(&dest)
	db.First(&dest)
	fmt.Printf("first: %+v \n", dest) // {Id:1 ConditionId:4}

	// sql: SELECT * FROM `policy` LIMIT 1
	// db.Table("policy").Take(&dest2)
	db.Take(&dest2)
	fmt.Printf("take: %+v \n", dest2) // {Id:1 ConditionId:4}

	// sql: SELECT * FROM `policy` ORDER BY `policy`.`id` DESC LIMIT 1
	// db.Table("policy").Last(&dest3)
	db.Last(&dest3)
	fmt.Printf("last: %+v \n", dest3) // {Id:3 ConditionId:3}
}


