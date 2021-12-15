package daytwo

import (
	"fmt"
	"math"
)

//int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
//uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型

func Func2() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
	//浮点型
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	//复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	//布尔值
	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	//字符串
	// len(str)、+或fmt.Sprintf、strings.Split、strings.contains、strings.HasPrefix、strings.HasSuffix
	// strings.Index()、strings.LastIndex()、strings.Join(a[]string, sep string)
	//类型转换
	//只有强制类型转换，没有隐式类型转换。T(表达式)
	var d, e = 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	f := int(math.Sqrt(float64(d*d + e*e)))
	fmt.Println(f)
}
