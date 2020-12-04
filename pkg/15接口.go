package main

import (
	"fmt"
	"math"
)

/*

*/

type geometry interface {
	area() float64
	perim() float64
}

// 在我们的例子中，我们将让 rect 和 circle 实现这个接口
type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
// 这里我们让 rect 实现了 geometry 接口。
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}


// circle 的实现。
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。
// 这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
func measure(g geometry) {
	fmt.Println("g:", g)
	fmt.Println("area:", g.area())
	fmt.Println("prim:", g.perim())
}

func main() {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}
	// 结构体类型 circle 和 rect 都实现了 geometry接口，
	// 所以我们可以使用它们的实例作为 measure 的参数。
	measure(r)
	measure(c)
}
