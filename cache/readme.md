## redis 连接池测试


```shell
127.0.0.1:6379> info clients
# Clients
connected_clients:1
client_recent_max_input_buffer:2
client_recent_max_output_buffer:0
blocked_clients:0
tracking_clients:0
clients_in_timeout_table:0
# 开始启动cache_test进行连接池测试, 连接池保持了10个连接
127.0.0.1:6379> info clients
# Clients
connected_clients:12
client_recent_max_input_buffer:2
client_recent_max_output_buffer:0
blocked_clients:0
tracking_clients:0
clients_in_timeout_table:0
127.0.0.1:6379>
```


```shell
127.0.0.1:6379[2]> MONITOR
OK
1603179046.338714 [0 172.26.0.1:37164] "SELECT" "2"
1603179046.340177 [2 172.26.0.1:37164] "set" "name" "go"

```