package create

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	mysql_driver "projects/example_demo/db/mysql"
	"testing"
)

type Resource struct {
	gorm.Model        // ID CreatedAt UpdatedAt DeletedAt
	AppId      int64  `gorm:"column:app_id;index"`
	Uin        string `gorm:"column:uin"` // 父账号uin
	CreatedBy  string `gorm:"column:created_by"`
	UpdatedBy  string `gorm:"column:updated_by"`
}

// 这里定义了多少字段，在create的时候，就会使用到哪些字段
type Policy struct {
	Resource
	UUID        string         `gorm:"column:uuid"`
	OriginID    string         `gorm:"column:origin_id"` // 告警1.0策略id
	Name        string         `gorm:"column:name"`
	Remark      string         `gorm:"column:remark"`       // 备注
	MonitorType string         `gorm:"column:monitor_type"` // 监控类型：MT_QCE=基础监控 MT_CUSTOM=自定义监控 MT_PROME=prometheus监控
	Namespace   sql.NullString `gorm:"column:namespace"`    // 策略类型：基础监控策略的viewName，自定义监控的namespace
	ConditionID uint           `gorm:"column:condition_id"` // 外键关联conditions表
	ProjectID   sql.NullInt32  `gorm:"column:project_id"`   // 项目id -1=无项目
	Enable      int32          `gorm:"column:enable"`       // 1=启用 0=停用
	IsDefault   int            `gorm:"column:is_default"`   // 是否为默认策略
	AlarmLevel  string         `gorm:"column:alarm_level"`  // 告警等级
}

func TestMiss(t *testing.T) {

	db := mysql_driver.BuildDB()
	policy := &Policy{
		Resource: Resource{
			AppId:     int64(1),
			Uin:       "1",
			CreatedBy: "1",
			UpdatedBy: "1",
		},
		ConditionID: uint(1),
		Enable:      int32(1),
		MonitorType: "11",
		Name:        "11",
		Namespace:   sql.NullString{String: "11", Valid: true},
		ProjectID:   sql.NullInt32{Int32: 11, Valid: true},
		Remark:      "11",
		UUID:       "11",
	}

	err := db.Create(policy).Error
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("success")
	}
}
