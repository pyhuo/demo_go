package main

import "fmt"

func main() {
	m := make([]int, 0)
	fmt.Printf("%v\n", m)
	fmt.Printf("len:%d\n", len(m))
	fmt.Printf("cap:%d\n", cap(m))
	m = append(m, 1)
	fmt.Printf("len:%d\n", len(m))
	fmt.Printf("cap:%d\n", cap(m))
}
