package main

import (
	"fmt"
	"os"
	"time"
)


func waitSig(c <-chan os.Signal, sig os.Signal, timeout time.Duration) {
	select {
	case s := <-c:
		if s != sig {
			fmt.Printf("signal was %v, want %v\n", s, sig)
		}
		fmt.Printf("signal ok was %v, want %v\n", s, sig)
	case <-time.After(timeout):
		fmt.Printf("timeout waiting for %v\n", sig)
	}
}

func Wait()  {
	c := make(chan os.Signal, 1)
	go func() {
		time.Sleep(time.Second*3)
		//signal.Notify(c, os.Interrupt)
		c <- os.Interrupt
		fmt.Println("send signal end")
	}()
	waitSig(c, os.Interrupt, time.Second*5)
	fmt.Println("main end")
}

func Select()  {
	c := make(chan os.Signal, 1)
	select {
	case s := <-c:
		fmt.Printf("signal ok was %v\n", s)
	default:
	}
	fmt.Printf("end")
}

func main()  {
	//Wait()
	Select()
}