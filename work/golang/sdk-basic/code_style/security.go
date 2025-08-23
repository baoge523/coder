package code_style

import "fmt"

// 在日常的日志输出，和json序列化传输中，其实我们是不希望将用户名或者密码在日志中打印出来

// 理解：日常日志输出的方式： print 、json序列化

/**
解决方式：
1、将username 和 password 设置成私有变量，提供get方法获取，但是以json装配的时候也会装配不上

2、提供String方法，里面不打印username和password， 但是不能控制json

3、了解print和json的处理方式，针对于json，可以装配到json对象中，但是输出为[]byte不能输出

*/

type DBConfig0 struct {
	username string `json:"username"`
	password string `json:"password"`
	Host     string `json:"supervisor_agent"`
	Port     int    `json:"port"`
	Ttl      int    `json:"ttl"`
	MinConn  int    `json:"min_conn"`
	MaxConn  int    `json:"max_conn"`
}

func (db *DBConfig0) getUsername() string {
	return db.username
}
func (db *DBConfig0) getPassword() string {
	return db.password
}

type DBConfig1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"supervisor_agent"`
	Port     int    `json:"port"`
	Ttl      int    `json:"ttl"`
	MinConn  int    `json:"min_conn"`
	MaxConn  int    `json:"max_conn"`
}

// 如果一个 `struct` 实现了 `String()` 方法，当你使用 `fmt.Println()` 或 `fmt.Printf()` 打印该 `struct` 时，会自动调用 `String()` 方法来获取其字符串表示。
func (db *DBConfig1) String() string {
	return fmt.Sprintf("supervisor_agent = %s, port = %d, ttl = %d, minConn = %d, maxConn = %d", db.Host, db.Port, db.Ttl, db.MinConn, db.MaxConn)
}

// ==================================

type DBConfig struct {
	Username UsernameValue `json:"username"`
	Password PasswordValue `json:"password"`
	Host     string        `json:"supervisor_agent"`
	Port     int           `json:"port"`
	Ttl      int           `json:"ttl"`
	MinConn  int           `json:"min_conn"`
	MaxConn  int           `json:"max_conn"`
}

// UsernameValue 不打印
type UsernameValue string

// PasswordValue 全部用*表示
type PasswordValue string

var AllowHidden = true

// Format 是print提供的拓展，可以用于对一个打印对象，做自定输出
// 注意点：对象方法，和指针方法的区别，可能会导致查询不到，需要根据应用场景来实现这个方法
// verb 是占位符，会根据打印情况变化的
func (text UsernameValue) Format(f fmt.State, verb rune) {
	fmt.Println(string(verb))
	if AllowHidden {
		fmt.Fprint(f, "")
	} else {
		fmt.Fprint(f, text)
	}
}

func (text PasswordValue) Format(f fmt.State, verb rune) {
	fmt.Fprint(f, "******")
}

// 如果一个 `struct` 实现了 `String()` 方法，当你使用 `fmt.Println()` 或 `fmt.Printf()` 打印该 `struct` 时，会自动调用 `String()` 方法来获取其字符串表示。
func (db *DBConfig) String() string {
	return fmt.Sprintf("username = %v, password = %s, supervisor_agent = %s, port = %d, ttl = %d, minConn = %d, maxConn = %d",
		db.Username, db.Password, db.Host, db.Port, db.Ttl, db.MinConn, db.MaxConn)
}

// MarshalJSON
func (text PasswordValue) MarshalJSON() ([]byte, error) {
	return []byte(`"******"`), nil
}
