package readyml

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

type Logger struct {
	Level        string `yaml:"info"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`
	LogInConsole bool   `yaml:"log-in-console"`
}

type System struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
	Env  string `yaml:"env"`
}

type Conf struct {
	Mysql  Mysql
	Logger Logger
	System System
}

var ConfigData = Conf{}

func ReadConf() {
	settingData, err := os.ReadFile("read-yml/setting.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(settingData, &ConfigData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", ConfigData)
}
