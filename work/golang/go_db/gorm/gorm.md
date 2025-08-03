## gorm
[官方文档](https://gorm.io/zh_CN/docs/query.html)


### golang在使用gorm时，实体类的默认值没有被更新到数据
在golang中，类型都会有默认值(也叫做零值)
接口类型或者指针类型、map、slice 等这些的默认值是 Null

| 类型        | 零值    | 零值描述  |
|-----------|-------|-------|
| string    | ""    | 空字符串  |
| int(数值类型) | 0     | 0值    |
| bool      | false | false |
| ...       | ...   | ...   |


**gorm中明确提示了:**
使用 struct 更新时, GORM 将只更新非零值字段

有时候，我们期望更新类型的零值到数据中、实体的定义信息如下：
```golang
// 通过该结构体更新时，如果期望将age、name的值改成其默认的零值，是不会更新的，GORM 将只更新非零值字段
type User struct {
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

// 我们可以将其改成如下的方式做更新
type User struct {
    Name sql.NullString `gorm:"column:name"`
    Age  sql.NullInt64    `gorm:"column:age"`
}
// 这里的sql是 go sdk中提供的，里面有多种Nullxxx的

// 当设置 Valid = true的时候，里面的string的默认值，也是可以被gorm正常更新的
//	if s.Valid {
//	   // use s.String
//	} else {
//	   // NULL value
//	}
type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}

```

