package main

import (
	"fmt"
	"time"
)

/*
1.go func
2.通信机制.
	* channel
	* 带缓冲channel
3.通道方向
4.通道关闭
5.select
6.超时实现
7.非阻塞选择器. select default
*/
func testBlockChannel()  {
	// 默认通道是无缓冲的，
	// 这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。
	// 可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
	message := make(chan string)
	go func() {
		fmt.Printf("go start..\n")
		// 阻塞写，直到主go程读
		message <- "ping"
		fmt.Printf("go end...\n")
		close(message)
	}()
	time.Sleep(time.Second)
	fmt.Printf("main go start.\n")
	fmt.Printf("msg:%s\n", <-message)
	_, ok := <-message
	if !ok {
		fmt.Printf("msg closed\n")
	}
	fmt.Printf("main go end.\n")
}

func testBufferChannel()  {
	// 通道缓冲区
	//通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
	ch := make(chan int, 10)

	go func() {
		fmt.Printf("go start..\n")
		// 非阻塞写.
		ch <- 10
		fmt.Printf("go wirite ..\n")
		ch <- 12
		fmt.Printf("go end...\n")
	}()

	time.Sleep(time.Second)
	fmt.Printf("main go start.\n")
	fmt.Printf("msg:%d\n", <-ch)
	fmt.Printf("msg:%d\n", <-ch)
	fmt.Printf("main go end.\n")
}

func testBufferChannelBlock()  {
	// 通道缓冲区
	//通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
	ch := make(chan int, 10)

	go func() {
		fmt.Printf("go start..\n")
		// 非阻塞写.
		for i:=0; i<10; i++{
			ch <- i
		}
		fmt.Printf("go wirite 0-10 ok\n")
		fmt.Printf("go wirite will block..\n")
		ch <- 11
		fmt.Printf("go end...\n")
	}()
	time.Sleep(time.Second * 5)
	fmt.Printf("main start.\n")
	//fmt.Printf("main first read:%d.\n", <-ch)
	//time.Sleep(time.Second * 1)
	for i:=0; i<11; i++ {
		fmt.Printf("i:%d main :%d.\n", i, <-ch)
	}
	fmt.Printf("main end.\n")
}


// 当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。
// 这个特性提升了程序的类型安全性。
func ping(pings chan <- string, msg string) {
	// ping 函数定义了一个只允许发送数据的通道。
	// 尝试使用这个通道来接收数据将会得到一个编译时错误。
	pings <- msg
}
func pong(pings <-chan string, pongs chan <- string) {
	// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
	msg := <-pings
	pongs <- msg
}
func testChannelReadWrite()  {
	//通道方向示例
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func testSelectChannel()  {
	// 在我们的例子中，我们将从两个通道中选择。
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()


	// 我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("i:%v received:%v\n", i, msg1)
		case msg2 := <-c2:
			fmt.Printf("i:%v received:%v\n", i, msg2)
		}
	}
	// 我们首先接收到值 "one"，然后就是预料中的 "two"了。
	// 注意从第一次和第二次 Sleeps 并发执行，总共仅运行了两秒左右。
}

func testSelectTimeOut()  {
	c1 := make(chan string, 1)
	// 在我们的例子中，假如我们执行一个外部调用，并在 2 秒后通过通道 c1 返回它的执行结果。
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 这里是使用 select 实现一个超时操作。res := <- c1 等待结果，<-Time.After 等待超时时间 1 秒后发送的值。
	// 由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout 1")
	}
}

func testSelectNoBlock()  {
	// 常规的通过通道发送和接收数据是阻塞的。
	// 然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
	messages := make(chan string)
	// 这里是一个非阻塞接收的例子。
	// 如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。
	// 如果不是，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	// 一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	signals := make(chan bool)
	// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	//go func() {
	//	signals <- true
	//}()
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func main() {
	//testBlockChannel()
	//testBufferChannel()
	//testBufferChannelBlock()
	//testChannelReadWrite()
	//testSelectChannel()
	//testSelectTimeOut()
	testSelectNoBlock()
}