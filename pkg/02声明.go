package main

import "fmt"

const PI = 3.14

func main() {
	var f = PI
	var c = (f - 32) * 5 / 9
	fmt.Printf("pi:%v f:%v c:%v\n", PI, f, c)
}