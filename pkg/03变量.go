package main
// https://books.studygolang.com/gopl-zh/ch2/ch2-03.html
/*
var 变量 类型

1.自动推断类型
2.自定初始化
3.变量组
var (
	x, y int				// x, y 自动初始化0
	a, s = 100, "abc"		// a 自动类型推断为int, s为字符串
)

4.简短模式
* 定义变量，同时初始化.
* 不能提供数据类型.
* 只能用在函数内部.
x := 100

*/


import "fmt"

func baseVar()  {
	// var 变量名字 类型 = 表达式
	var w string
	w = "hello world"
	fmt.Println(w)

	var i, j, k int                 // int, int, int
	var b, f, s = true, 2.3, "four" // bool, float64, string
	fmt.Printf("i:%v j:%v k:%v\n", i, j, k)
	fmt.Printf("b:%v f:%v s:%v\n", b, f, s)
}

func shortDefVar()  {
	// 简短变量声明
	i, j := 1, 0
	fmt.Printf("i:%v j:%v\n", i, j)

	var pi float64 = 3.14
	var names []string
	var err error
	fmt.Printf("pi:%v\n", pi)
	fmt.Printf("err:%v\n", err)
	fmt.Printf("names:%v\n", names)
}

func ptrVar()  {
	//2.3.2. 指针
	x := 1
	p := &x
	fmt.Printf("addr:%v v:%v\n", &x, x)
	fmt.Printf("addr:%v p:%v v:%v\n", &p, p, *p)
}
func returnPtr() * int{
	// 返回函数中局部变量的地址也是安全的
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func fReturnLocal()  {
	//调用returnPtr函数时创建局部变量v，
	//在局部变量地址被返回之后依然有效，
	//因为指针p依然引用这个变量
	var p = returnPtr()
	fmt.Printf("ptr:%v v:%v\n", p, *p)
	var x, y = 0, 0
	y = incr(&x)
	fmt.Printf("x:%v y:%v\n", x, y)
	y = incr(&x)
	fmt.Printf("x:%v y:%v\n", x, y)
}

func main() {
	//baseVar()
	//shortDefVar()
	//ptrVar()
	//fReturnLocal()
}