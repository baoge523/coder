package main

import "fmt"
import _ "test/usercase/syst/init/pack_a"
import _ "test/usercase/syst/init/pack_b"

func init()  {
	fmt.Println("main init")
}


func main() {
	fmt.Println("start main")
}
