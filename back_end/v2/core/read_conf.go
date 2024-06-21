package core

import (
	"database_lesson/config"
	"database_lesson/global"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{} //实例的创建和指向
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal:%v", err)
	}
	log.Println("config init success")
	global.Config = c
	//fmt.Println(c)
}
