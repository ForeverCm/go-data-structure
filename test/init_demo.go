package test

import (
"fmt"
)

var ee int
var dd int64 = func() int64{
	fmt.Println("calling cc() in main.go")
	return 1
}()

func init() {
	fmt.Println("main.init1")
}
func init() {
	fmt.Println("main. init2")
}

//func cc() int64 {
//	fmt.Println("calling cc() in main.go")
//	return 1
//}
func main() {
	fmt.Println("main.ing")
	fmt.Println(dd)
}
