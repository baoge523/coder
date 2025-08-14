package standard_lib

type User struct {
	Name string
	Age  int
}

type Runner interface {
	Run()
}

func (u *User) Run() {
	
}
