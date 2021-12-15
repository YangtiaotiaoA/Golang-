package daytwo

import "fmt"

//变量初始化
var name string
var age int
var isOk bool
var (
	a string
	b int
	c bool
	d float32
)
var name1 string = "yanglei"
var age1 int = 24
var name2, age2 = "yanglie", 24

//类型推导
var name3 = "yanglei"
var age3 = 18

func foo() (int, string) {
	return 24, "yanglei"
}
func Func1() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//在函数内部，可以使用:= 方式声明并初始化变量
	name4 := "yanglei"
	fmt.Println(name4)
	//匿名变量
	name5, _ := foo()
	fmt.Println("name=", name5)
	_, age5 := foo()
	fmt.Println("age=", age5)
	//常量
	const (
		pi = 3.1415
		e  = 2.7182
	)
	//iota是go语言的常量计数器，只能在常量的表达式中使用
	//iota在const关键字出现时将被重置为0,在定义枚举时很有用。
	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)
}
