package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string

	EtcdHost string
	EtcdPort string

	UserServiceAddress     string
	TaskServiceAddress     string
	EmployeeServiceAddress string
	AuthServiceAddress     string
	ProducerServiceAddress string
	ProductServiceAddress  string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDbName   int
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadServer(file)
	LoadRedisData(file)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
	EmployeeServiceAddress = file.Section("server").Key("EmployeeServiceAddress").String()
	AuthServiceAddress = file.Section("server").Key("AuthServiceAddress").String()
	ProducerServiceAddress = file.Section("server").Key("ProducerServiceAddress").String()
	ProductServiceAddress = file.Section("server").Key("ProductServiceAddress").String()
}

func LoadRedisData(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
}
