package main

import "fmt"

/*
函数
1.变参函数
2.多返回值
3.多返回值函数
4.闭包函数
4.递归
*/

// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
	return a + b
}

func plusPlus(a, b, c int)  int{

	return a + b +c
}

// 变参函数
func sum(nums ...int)  int{
	ret := 0
	for _, v := range nums {
		ret += v
	}
	return ret
}

// 多返回值函数
func mulReturn() (int, int) {
	return 3, 7
}

// 闭包
// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。
// 这个返回的函数使用闭包的方式 隐藏 变量 i。
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 递归
// face 函数在到达 face(0) 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}


func main() {
	seqF := intSeq()
	for i:=0; i<3; i++ {
		fmt.Printf("i:%v v:%v\n", i, seqF())
	}
	fmt.Printf("4!=%v\n", fact(4))
	//fmt.Printf("5!=%v\n", fact(5))
	//fmt.Printf("6!=%v\n", fact(6))
	//fmt.Printf("7!=%v\n", fact(7))
	// 4 * f(3)
	// 4 * 3 * f(2)
	// 4 * 3 * 2 * f(1)
}