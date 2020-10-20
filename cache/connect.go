package cache

import (
	"demo_go/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	//定义常量
	RedisClient *redis.Pool //redis 连接池
	//REDIS_DB    int
)

func init() {
	// 从配置文件获取redis的ip以及db
	RedisClient = newPool()
	infoClients()
}

func infoClients()  {
	fmt.Printf("RedisClient:%v\n", RedisClient)
}

//初始化连接池
func newPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle:     5,                 //最初的连接数量
		//MaxActive:   0,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		MaxActive:   5,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //超时时间 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(config.AppConf.RedisConf["type"], config.AppConf.RedisConf["address"])
			if err != nil {
				return nil, err
			}
			if _, errAuth := con.Do("auth", config.AppConf.RedisConf["auth"]); err != nil {
				return nil, errAuth
			}
			return con, nil
		},
	}
}

func info() {
	rc := RedisClient.Get()
	//用完连接放回到连接池
	defer rc.Close()
	r, e := redis.String(rc.Do("info", "clients"))
	if e != nil {
		panic(e)
	}
	fmt.Println("#######")
	fmt.Printf("infos:\n%v\n", r)
	fmt.Println("#######")
}