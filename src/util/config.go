package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/** USAGE:
config, err := util.GetConfig()
if err != nil {
	log.Printf("错了2: #%v", err)
}
fmt.Println(config.Mysql)
*/
var configFile []byte

var Configs = Config{}

type Config struct {
	DisablePathCorrection             bool   `yaml:"DisablePathCorrection"`
	EnablePathEscape                  bool   `yaml:"EnablePathEscape"`
	FireMethodNotAllowed              bool   `yaml:"FireMethodNotAllowed"`
	DisableBodyConsumptionOnUnmarshal bool   `yaml:"DisableBodyConsumptionOnUnmarshal"`
	TimeFormat                        string `yaml:"TimeFormat"`
	Charset                           string `yaml:"Charset"`
	AppPort                           string `yaml:"AppPort"`
	Mysql                             Mysql  `yaml:"Mysql"`
}

type Mysql struct {
	Port   string `yaml:"Port"`
	Name   string `yaml:"Name"`
	Pwd    string `yaml:"Pwd"`
	Ip     string `yaml:"Ip"`
	DBName string `yaml:"DBName"`
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("src/config/iris.yml")
	if err != nil {
		panic("Sys config read err")
	}
	err = yaml.Unmarshal(configFile, &Configs)
	if err != nil {
		panic(err)
	}
}
