package main

import "fmt"

// DataProcessor 数据处理器接口
type DataProcessor interface {
	Connect() string
	ExtractData() string
	ProcessData(data string) string
	SaveData(data string) string
}

// Template 模板类
type Template struct {
	processor DataProcessor
}

func NewTemplate(processor DataProcessor) *Template {
	return &Template{processor: processor}
}

// Process 模板方法（定义算法骨架）
func (t *Template) Process() {
	fmt.Println("=== 开始数据处理流程 ===")
	
	// 步骤1: 连接数据源
	fmt.Println("\n步骤1: 连接数据源")
	result := t.processor.Connect()
	fmt.Println(result)
	
	// 步骤2: 提取数据
	fmt.Println("\n步骤2: 提取数据")
	data := t.processor.ExtractData()
	fmt.Println(data)
	
	// 步骤3: 处理数据
	fmt.Println("\n步骤3: 处理数据")
	processedData := t.processor.ProcessData(data)
	fmt.Println(processedData)
	
	// 步骤4: 保存数据
	fmt.Println("\n步骤4: 保存数据")
	saveResult := t.processor.SaveData(processedData)
	fmt.Println(saveResult)
	
	fmt.Println("\n=== 数据处理流程完成 ===\n")
}

// MySQLProcessor MySQL数据处理器
type MySQLProcessor struct{}

func (m *MySQLProcessor) Connect() string {
	return "连接到 MySQL 数据库"
}

func (m *MySQLProcessor) ExtractData() string {
	return "从 MySQL 提取数据: [用户数据]"
}

func (m *MySQLProcessor) ProcessData(data string) string {
	return fmt.Sprintf("MySQL处理: %s -> [清洗后的用户数据]", data)
}

func (m *MySQLProcessor) SaveData(data string) string {
	return fmt.Sprintf("保存到 MySQL: %s", data)
}

// CSVProcessor CSV文件处理器
type CSVProcessor struct{}

func (c *CSVProcessor) Connect() string {
	return "打开 CSV 文件"
}

func (c *CSVProcessor) ExtractData() string {
	return "从 CSV 读取数据: [销售数据]"
}

func (c *CSVProcessor) ProcessData(data string) string {
	return fmt.Sprintf("CSV处理: %s -> [格式化的销售数据]", data)
}

func (c *CSVProcessor) SaveData(data string) string {
	return fmt.Sprintf("保存到 CSV: %s", data)
}

// APIProcessor API数据处理器
type APIProcessor struct{}

func (a *APIProcessor) Connect() string {
	return "连接到 REST API"
}

func (a *APIProcessor) ExtractData() string {
	return "从 API 获取数据: [订单数据]"
}

func (a *APIProcessor) ProcessData(data string) string {
	return fmt.Sprintf("API处理: %s -> [标准化的订单数据]", data)
}

func (a *APIProcessor) SaveData(data string) string {
	return fmt.Sprintf("通过 API 保存: %s", data)
}

func main() {
	fmt.Println("=== 模板方法模式 - 数据处理系统 ===\n")
	
	// 处理 MySQL 数据
	fmt.Println("【MySQL 数据处理】")
	mysqlProcessor := &MySQLProcessor{}
	mysqlTemplate := NewTemplate(mysqlProcessor)
	mysqlTemplate.Process()
	
	// 处理 CSV 数据
	fmt.Println("\n【CSV 数据处理】")
	csvProcessor := &CSVProcessor{}
	csvTemplate := NewTemplate(csvProcessor)
	csvTemplate.Process()
	
	// 处理 API 数据
	fmt.Println("\n【API 数据处理】")
	apiProcessor := &APIProcessor{}
	apiTemplate := NewTemplate(apiProcessor)
	apiTemplate.Process()
}
