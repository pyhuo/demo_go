## go配置mysql连接池
* [mysql配置文档](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_connections)
* max-connections=10000

## 1.ab 发起并发请求测试
* -c: 并发量100
* -n: 总请求量300
```shell
➜  ~ ab -c 100 -n 300 "http://localhost:8080/pool"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Finished 300 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /pool
Document Length:        304 bytes

Concurrency Level:      100
Time taken for tests:   4.073 seconds
Complete requests:      300
Failed requests:        0
Total transferred:      128400 bytes
HTML transferred:       91200 bytes
Requests per second:    73.66 [#/sec] (mean)
Time per request:       1357.601 [ms] (mean)
Time per request:       13.576 [ms] (mean, across all concurrent requests)
Transfer rate:          30.79 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   2.2      1       8
Processing:  1001 1012   8.2   1010    1036
Waiting:     1001 1012   8.1   1010    1036
Total:       1002 1014   9.5   1012    1043

Percentage of the requests served within a certain time (ms)
  50%   1012
  66%   1016
  75%   1019
  80%   1021
  90%   1031
  95%   1035
  98%   1039
  99%   1040
 100%   1043 (longest request)
➜  ~
```

## 2.PoolHandler 中触发sql查询
* 设置空闲连接为10个
* 最大可以打开的连接为10个
mysql.conf中可以修改最大连接数.

db/db.go
```go
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
```
controller/pool.go
```go
func PoolHandler(c *gin.Context) {
	time.Sleep(time.Second)
	var user model.User
	tx := db.DB.Raw("SELECT id, name, age FROM users WHERE name = ? limit 1", "D42").Scan(&user)
	if tx.Error != nil {
		panic(tx.Error)
	}
	logger.Debugf("raw sql id:%v name:%v age:%v", user.ID, user.Name, user.Age)
	c.JSON(200, user)
}
```

## 3.查看测试过程中请求mysql创建的请求

```sql
SHOW PROCESSLIST;
```
```
Id	User	Host	db	Command	Time	State	Info
8	root	localhost:51148	demo_go	Query	0	starting	SHOW PROCESSLIST
900	root	localhost:64598	demo_go	Sleep	0		
901	root	localhost:64754	demo_go	Sleep	0		
902	root	localhost:64755	demo_go	Sleep	0		
903	root	localhost:64756	demo_go	Sleep	0		
904	root	localhost:64757	demo_go	Sleep	0		
905	root	localhost:64758	demo_go	Sleep	0		
906	root	localhost:64759	demo_go	Sleep	0		
907	root	localhost:64760	demo_go	Sleep	0		
908	root	localhost:64761	demo_go	Sleep	0		
909	root	localhost:64762	demo_go	Sleep	0		
```