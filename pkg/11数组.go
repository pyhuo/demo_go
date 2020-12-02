package main

import "fmt"

/*
1.数组是一个由固定长度的特定类型元素组成的序列，长度固定.
	* len: 固定

2.Slice: 动态

初始化:
1.默认初始化为每个元素都被初始化为元素类型对应的零值
2. ... 省略号表示，则表示数组的长度是根据初始化值的个数来计算

操作:
1.索引数组.
2.for ... range 遍历数组
*/

var a [3]int

func main() {
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Printf("%v\n", a)

	for i, v := range a {
		fmt.Printf("i:%d v:%d arr:%v\n", i, v, a[i])
	}

	var q [3]string = [3]string{"1", "0", "c"}
	fmt.Printf("init q len:%v v:%v\n", len(q), q)

	var p [3]string = [...]string{"1", "0", "c"}
	fmt.Printf("init p len:%v v:%v\n", len(p), p)
}