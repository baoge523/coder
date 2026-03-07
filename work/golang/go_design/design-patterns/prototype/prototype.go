package main

import "fmt"

// Cloneable 可克隆接口
type Cloneable interface {
	Clone() Cloneable
}

// WorkExperience 工作经历
type WorkExperience struct {
	Company  string
	Position string
	Years    int
}

func (w *WorkExperience) Clone() *WorkExperience {
	return &WorkExperience{
		Company:  w.Company,
		Position: w.Position,
		Years:    w.Years,
	}
}

// Resume 简历
type Resume struct {
	Name       string
	Age        int
	Experience *WorkExperience
}

// ShallowClone 浅拷贝
func (r *Resume) ShallowClone() *Resume {
	return &Resume{
		Name:       r.Name,
		Age:        r.Age,
		Experience: r.Experience, // 引用相同的对象
	}
}

// DeepClone 深拷贝
func (r *Resume) DeepClone() *Resume {
	return &Resume{
		Name:       r.Name,
		Age:        r.Age,
		Experience: r.Experience.Clone(), // 克隆新对象
	}
}

func (r *Resume) Display() {
	fmt.Printf("姓名: %s, 年龄: %d\n", r.Name, r.Age)
	if r.Experience != nil {
		fmt.Printf("工作经历: %s - %s (%d年)\n",
			r.Experience.Company,
			r.Experience.Position,
			r.Experience.Years)
	}
}

// 实际应用示例：配置对象克隆

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Options  map[string]string
}

func (d *DatabaseConfig) Clone() *DatabaseConfig {
	// 深拷贝 map
	optionsCopy := make(map[string]string)
	for k, v := range d.Options {
		optionsCopy[k] = v
	}
	
	return &DatabaseConfig{
		Host:     d.Host,
		Port:     d.Port,
		Username: d.Username,
		Password: d.Password,
		Database: d.Database,
		Options:  optionsCopy,
	}
}

func (d *DatabaseConfig) Display() {
	fmt.Printf("数据库配置: %s:%d/%s (用户: %s)\n",
		d.Host, d.Port, d.Database, d.Username)
	fmt.Printf("选项: %v\n", d.Options)
}

func main() {
	fmt.Println("=== 原型模式 ===\n")
	
	// 简历克隆示例
	fmt.Println("【简历克隆示例】")
	original := &Resume{
		Name: "张三",
		Age:  28,
		Experience: &WorkExperience{
			Company:  "腾讯",
			Position: "高级工程师",
			Years:    3,
		},
	}
	
	fmt.Println("\n原始简历:")
	original.Display()
	
	// 浅拷贝
	fmt.Println("\n浅拷贝:")
	shallowCopy := original.ShallowClone()
	shallowCopy.Name = "李四"
	shallowCopy.Experience.Company = "阿里巴巴" // 会影响原始对象
	
	fmt.Println("修改后的浅拷贝:")
	shallowCopy.Display()
	fmt.Println("原始简历（被影响）:")
	original.Display()
	
	// 深拷贝
	fmt.Println("\n深拷贝:")
	original.Experience.Company = "腾讯" // 恢复原值
	deepCopy := original.DeepClone()
	deepCopy.Name = "王五"
	deepCopy.Experience.Company = "字节跳动" // 不会影响原始对象
	
	fmt.Println("修改后的深拷贝:")
	deepCopy.Display()
	fmt.Println("原始简历（未被影响）:")
	original.Display()
	
	// 配置克隆示例
	fmt.Println("\n\n【配置克隆示例】")
	baseConfig := &DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "password",
		Database: "test_db",
		Options: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "true",
		},
	}
	
	fmt.Println("\n基础配置:")
	baseConfig.Display()
	
	// 克隆用于生产环境
	prodConfig := baseConfig.Clone()
	prodConfig.Host = "prod.example.com"
	prodConfig.Database = "prod_db"
	prodConfig.Options["maxConnections"] = "100"
	
	fmt.Println("\n生产环境配置:")
	prodConfig.Display()
	
	fmt.Println("\n基础配置（未被影响）:")
	baseConfig.Display()
}
