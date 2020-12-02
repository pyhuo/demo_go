package main

import "fmt"

// 创建变量, new(T), 分配内存，变量在堆中
// 初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T

func newInt()  {
	i, j := 0, 0
	fmt.Printf("i:%p j:%p\n", &i, &j)
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Printf("new p addr:%p p:%v v:%d\n", &p, p, *p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
}

func main() {
	newInt()
}