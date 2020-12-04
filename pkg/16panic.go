package main

import (
	"fmt"
	"runtime"
	"time"
)

func handlerPanic()  {
	if p := recover(); p != nil {
		err := fmt.Errorf("%v", p)
		fmt.Println(err)
	}
}

func main() {
	go func() {
		defer handlerPanic()
		fmt.Printf("in goroutine\n")
		panic("发生异常 in goroutine")
		// 不会被执行，只能恢复后，返回到主go程
		fmt.Printf("in goroutine after panic")
	}()
	time.Sleep(1 * time.Second)
	fmt.Printf("in main\n")
	defer handlerPanic()
	panic("发送异常 in main")
	// 不会执行
	fmt.Printf("in main after panic")
	runtime.Error()
}

/*
goroutine:
T:0		g1:main			g2:func
t1: 	sleep1			 fmt.
t2:					 	 panic 终止
						 recover: 必须在对应的panic go程中
1s:		fmt
*/

/*
1.panic 发生之前，我们需要为当前的go程注册defer recover恢复.
*/