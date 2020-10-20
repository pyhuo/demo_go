package cache

import (
	"demo_go/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"testing"
	"time"
)

func init()  {
	config.Init("prod", "../conf/app.yml")
}

func readName()  {
	rc := RedisClient.Get()
	defer rc.Close()
	r, e := redis.String(rc.Do("get", "name"))
	fmt.Println("name:", r, e)
}

func TestCacheSet(t *testing.T)  {
	rc := RedisClient.Get()
	//用完连接放回到连接池
	defer rc.Close()
	//time.Sleep(time.Second * 10)
	r, e := redis.String(rc.Do("set", "name", "go"))
	fmt.Println("set", r, e)
}

func TestCacheRead(t *testing.T) {
	//测试每一个go程获取一个redis连接
	var wg sync.WaitGroup
	info()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			readName()
			time.Sleep(time.Second * 1)
			wg.Done()
		}()
	}
	info()
	wg.Wait()
	// 开启四个go程，刚好占用五个连接
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			readName()
			time.Sleep(time.Second * 1)
			wg.Done()
		}()
	}
	info()
	wg.Wait()
}