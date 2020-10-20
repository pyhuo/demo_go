package cache

import (
	"demo_go/config"
	"demo_go/logger"
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
	logger.Debugf("name:%v err:%v\n", r, e)
}

func TestCacheSet(t *testing.T)  {
	rc := RedisClient.Get()
	//用完连接放回到连接池
	defer rc.Close()
	//time.Sleep(time.Second * 10)
	r, e := redis.String(rc.Do("set", "name", "go"))
	logger.Debug("set", r, e)
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

func TestCacheSetJson(t *testing.T)  {
	m := make(map[string]string, 10)
	m["usename"] = "666"
	m["phoneNum"] = "8888"

	rc := RedisClient.Get()
	defer rc.Close()
	k := "user"
	_, e := HandlerJson(rc, k, m)
	logger.Debugf("e:%v", e)
	user := JsonToBean(rc, k)
	logger.Debugf("user from cache:%v\n", user)
}