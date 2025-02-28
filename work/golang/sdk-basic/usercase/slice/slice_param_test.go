package slice

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
验证切片对象 和 切片指针对象的参数传递

该两种方式都能改变其内部的值，因为切片本身就是一个指针类型

但是如果在遍历切片对象时，如果使用foreach时，使用的是临时变量，会导致修改无效

*/

type User struct {
	Name string
	AA   *AA
}
type AA struct {
	A string
}

type User1 struct {
	Name string
	AA   AA1
}
type AA1 struct {
	A string
}

/*
*

	遍历切片对象时，使用index方式访问切片中的数据，修改成功
*/
func TestParamSlice(t *testing.T) {
	users2 := make([]User, 0, 0)

	users2 = append(users2, User{Name: "lisi", AA: &AA{"aa"}})

	m := make(map[string]*User)
	m["11"] = &User{Name: "lisi222", AA: &AA{"aa222"}}
	for i, _ := range users2 {
		u := m["11"]
		users2[i].Name = u.Name
		users2[i].AA = u.AA
	}
	for _, user := range users2 {
		fmt.Printf("%s, %+v \n", user.Name, user.AA)
	}
}

/*
遍历切片对象时，使用foreach方式访问切片中的数据，因为是临时变量，会导致修改失败
*/
func TestParamSlice2(t *testing.T) {
	users2 := make([]User, 0, 0)

	users2 = append(users2, User{Name: "lisi", AA: &AA{"aa"}})

	m := make(map[string]*User)
	m["11"] = &User{Name: "lisi222", AA: &AA{"aa222"}}
	// user是临时变量，如果 users2是struct切片，那么修改不成功
	for _, user := range users2 {
		u := m["11"]
		user.Name = u.Name
		user.AA = u.AA
	}
	for _, user := range users2 {
		fmt.Printf("%s, %+v \n", user.Name, user.AA)
	}
}

func TestJson(t *testing.T) {

	u := User{Name: "lisi", AA: &AA{"aa"}}
	marshal, _ := json.Marshal(u)
	fmt.Println(string(marshal))
	var u1 User
	json.Unmarshal(marshal, &u1)
	fmt.Println(u1)
}
