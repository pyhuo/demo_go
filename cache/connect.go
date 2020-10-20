package cache

import (
	"demo_go/config"
	"demo_go/logger"
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
	logger.Debugf("RedisClient:%v\n", RedisClient)
}

//初始化连接池
func newPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle:     5,                 //最初的连接数量
		//MaxActive:   5,               //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		MaxActive:   0,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //超时时间 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(config.AppConf.RedisConf["type"], config.AppConf.RedisConf["address"])
			if err != nil {
				return nil, err
			}
			if _, errAuth := con.Do("auth", config.AppConf.RedisConf["auth"]); err != nil {
				return nil, errAuth
			}
			if _, err := con.Do("SELECT", config.AppConf.RedisConf["db"]); err != nil {
				_ = con.Close()
				return nil, err
			}
			return con, nil
		},
	}
}

