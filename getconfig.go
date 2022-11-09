package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	CasLoinTicketName = "CasLoinTicket"
)

var (
	configData configUtils
)

type configUtils struct {
	AllowDomain   []string `yaml:"allowDomain"`
	AdminUserInfo User     `yaml:"adminUserInfo"`
}

type User struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

func init() {
	loadYamlConfig()
	initRedis()
}

// 记载程序运行时需要读取的配置信息
func loadYamlConfig() {
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		// do some thing
		// like panic
	}
	err = yaml.Unmarshal(file, &configData)
	if err != nil {
		// do some thing
		// like panic
	}
}
