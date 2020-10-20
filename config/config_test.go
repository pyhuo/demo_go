package config

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"testing"
)

func TestNewConf(t *testing.T) {

	//mysql配置文件信息
	var MySqlConfig = map[string]string{
		"mysql":     "mysql",
		"name":      "root",
		"password":  "123456",
		"type":      "tcp",
		"host":      "3306",
		"allconfig": "root:@tcp(localhost:3306)/mydb",
	}
	var RedisConfig = map[string]string{
		"name":    "redis",
		"type":    "tcp",
		"address": "127.0.0.1:6379",
		"auth":    "123456",
		"db":		"2",
	}

	var LogConfig = map[string]string{
		"default_path":    "./log/log.debug",
	}

	conf := AppConfInfo{
		AppName:         "bdev",
		HttpPort:        5008,
		GrpcListen:      ":5008",
		RunMode:         "prod",
		AutoRender:      false,
		CopyRequestBody: true,
		EnableDocs:      true,
		OrmDebug:        true,
		PgDataSource:    "user=postgres password=postgres dbname=test host=127.0.0.1 port=5432 sslmode=disable",
		LogLevel:        logs.LevelDebug,
		JwtSalt:         "testsalt",
		RedisConf: RedisConfig,
		MysqlConf: MySqlConfig,
		LogConf:  LogConfig,
	}


	data, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", string(data))
	if err := iot.WriteFile("../conf/app_template.yml", data, 0644); err != nil {
		t.Fatal(err)
	}

}

func TestLoadConf(t *testing.T) {
	data, err := LoadConf("../conf/app.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}

func TestLoadConfFromData(t *testing.T) {
	data, err := LoadConfFromData(TestData, "yml")
	if err != nil {
		t.Logf("data:%v", data)
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}
