package cache

import (
	"demo_go/config"
	"demo_go/logger"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	//"time"
)

//封装业务层的缓存
func info() {
	rc := RedisClient.Get()
	defer rc.Close()
	r, e := redis.String(rc.Do("info", "clients"))
	if e != nil {
		panic(e)
	}
	logger.Debug("#######")
	logger.Debugf("infos:\n%v\n", r)
	logger.Debug("#######")
}


//操作json
func HandlerJson(con redis.Conn, k string, v interface{}) (reply interface{}, err error){
	data, e := json.Marshal(v)
	config.ErrCheck(e)
	return con.Do("set", k, data)
}

//将读取到的json转为bean
func JsonToBean(con redis.Conn, k string) map[string]interface{} {
	m := make(map[string]interface{})
	data, e := redis.Bytes(con.Do("get", k))
	config.ErrCheck(e)
	err := json.Unmarshal(data, &m)
	config.ErrCheck(err)
	return m
}


