package main

import "fmt"

/*
https://github.com/pyhuo/gostudy/blob/master/doc/09.golang%E6%8C%87%E9%92%88%E7%BB%93%E6%9E%84%E4%BD%93%E6%8E%A5%E5%8F%A3.md
Go 指针
指针使用流程：
1.定义指针变量。
2.为指针变量赋值。
3.访问指针变量中指向地址的值。
*/

type Person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

// 这里的 area 方法有一个接收器类型 rect。
func (r *rect) area() int{
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。这里是一个值类型接收器的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}




func main() {
	var ip *int        /* 指向整型: nil*/
	var fp *float32    /* 指向浮点型: nil */
	tInt := 1
	tFloat := float32(1.4)
	ip = new(int)
	fp = new(float32)
	*ip = tInt
	*fp = tFloat
	fmt.Printf("ip:%v\n", *ip)
	fmt.Printf("fp:%v\n", *fp)


	// 使用这个语法创建了一个新的结构体元素。
	fmt.Println(Person{"Bob", 20})

	// 使用点来访问结构体字段。
	s := Person{name: "Sean", age: 50}
	fmt.Printf("name:%v age:%v\n", s.name, s.age)

	r := rect{width: 10, height: 5}
	// 这里我们调用上面为结构体定义的两个方法。
	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}