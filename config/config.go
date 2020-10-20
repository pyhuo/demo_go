package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"os"
	"strings"
)

type AppConfInfo struct {
	AppName         string `yml:"app_name" json:"app_name"`
	HttpPort        int    `yml:"http_port" json:"http_port"`
	RunMode         string `yml:"run_mode" json:"run_mode"`
	AutoRender      bool   `yml:"auto_render" json:"auto_render"`
	CopyRequestBody bool   `yml:"copy_request_body" json:"copy_request_body"`
	GrpcListen      string `yml:"grpc_listen" json:"grpc_listen"`
	PgDataSource    string `yml:"pg_data_source" json:"pg_data_source"`
	OrmDebug        bool   `yml:"orm_debug" json:"orm_debug"`
	EnableDocs      bool   `yml:"enable_docs" json:"enable_docs"`
	LogLevel        int    `yml:"log_level" json:"log_level"`
	JwtSalt         string `yaml:"jwt_salt" json:"jwt_salt"`
	RedisConf 	map[string]string	`yml:"redis_conf" json:"redis_conf"`
	MysqlConf 	map[string]string	`yml:"mysql_conf" json:"redis_conf"`
	LogConf 	map[string]string	`yml:"log_conf" json:"log_conf"`
}

var AppConf *AppConfInfo

func init() {
	//Init("debug", "")
}

var TestData = []byte(`
appname: bdev
httpport: 5008
runmode: prod
autorender: false
copyrequestbody: true
grpclisten: :5008
pgdatasource: user=postgres password=postgres dbname=test host=127.0.0.1 port=5432
  sslmode=disable
ormdebug: true
enabledocs: true
loglevel: 7
jwt_salt: testsalt
redisconf:
  address: 127.0.0.1:6379
  auth: ""
  name: redis
  type: tcp
mysqlconf:
  allconfig: root:@tcp(localhost:3306)/mydb
  host: "3306"
  mysql: mysql
  name: root
  password: ""
  type: tcp
`)

func Init(env string, filepath string) {
	var err error
	if AppConf != nil {
	}
	if filepath == "" {
		filepath = "conf/app.yml"
	}
	if env == "debug" {
		AppConf, err = LoadConfFromData(TestData, "yml")
	}else{
		AppConf, err = LoadConf(filepath)
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("env:%v err:%v\n\n", env, err)
}

func LoadConf(filepath string) (item *AppConfInfo, err error) {
	if filepath == "" {
		return nil, errors.New("filepath is empty, must use --config xxx.yml/json")
	}

	data, err := iot.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	item = &AppConfInfo{}
	if strings.HasSuffix(filepath, ".json") {
		err = json.Unmarshal(data, item)
	} else if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
		err = yaml.Unmarshal(data, item)
	} else {
		return nil, errors.New("you config file must be json/yml")
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

func LoadConfFromData(data []byte, t string) (item *AppConfInfo, err error) {
	// data 可以是json/yml格式的数据. 调用方需要指定t.
	item = &AppConfInfo{}
	if t == "json" {
		err = json.Unmarshal(data, item)
	} else if t == "yml" {
		err = yaml.Unmarshal(data, item)
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}

func ErrCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}

