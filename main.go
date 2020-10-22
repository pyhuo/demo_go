package main

import (
	"demo_go/logger"
	_ "demo_go/model"
	"demo_go/web"
	_ "demo_go/web"
	"sync"
)

func init() {
	// 启动http server
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Debug("ginStart server start ")
		web.GinStart()
		wg.Done()
	}()
	wg.Wait()
	logger.Debug("ser done")
}

