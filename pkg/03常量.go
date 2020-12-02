package main

import (
	"fmt"
)

/*
常量:
1.在编译器分配，没有对应地址.

2.常量生成器

3.无类型常量
*/


type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

	Pi = 3.312321
)


func main() {
	fmt.Printf("Sunday:%v\n", Sunday)
	fmt.Printf("Monday:%v\n", Monday)
	fmt.Printf("Tuesday:%v\n", Tuesday)
	fmt.Printf("Wednesday:%v\n", Wednesday)
	fmt.Printf("Thursday:%v\n", Thursday)
	fmt.Printf("Friday:%v\n", Friday)
	fmt.Printf("Saturday:%v\n", Saturday)
	fmt.Printf("Pi:%v\n", Pi)
}