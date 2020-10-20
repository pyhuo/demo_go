package main

import (
	"demo_go/config"
	"fmt"
)

//初始化连接池
func init() {
	config.Init("prod", "./conf/app.yml")
}

func main() {
	fmt.Println("happy go")
}
