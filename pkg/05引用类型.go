package main

import "fmt"

/*
引用类型: slice, map, channel 三种预定义类型

内置函数new按指定类型长度分配零值内存，返回指针，并不关心类型内部构造和初始化方式。
而引用类型则必须使用make函数创建，
编译器会将make转换为目标类型专用的创建函数（或指令），以确保完成全部内存分配和相关属性初始化

*/

func mkslice(size int)  [] int{
	s := make([]int, 0, size)
	s = append(s, 100)
	return s
}

func mkmap(size int)  map[string]int{
	m := make(map[string]int)
	m["key"] = 1
	return m
}


func main() {
	s := mkslice(10)
	fmt.Printf("s:%v len:%v cap:%v\n", s, len(s), cap(s))
}

