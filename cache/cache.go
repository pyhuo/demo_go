package cache

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
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
	fmt.Println("#######")
	fmt.Printf("infos:\n%v\n", r)
	fmt.Println("#######")
}


//操作json
func HandlerJson(con redis.Conn, k string, v interface{}) (reply interface{}, err error){
	data, e := json.Marshal(v)
	errCheck(e)
	return con.Do("set", k, data)
}

//将读取到的json转为bean
func JsonToBean(con redis.Conn, k string) map[string]interface{} {
	m := make(map[string]interface{})
	data, e := redis.Bytes(con.Do("get", k))
	errCheck(e)
	err := json.Unmarshal(data, &m)
	errCheck(err)
	return m
}


func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}