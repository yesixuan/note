package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/** USAGE:
config, err := util.GetConfig()
if err != nil {
	log.Printf("错了2: #%v", err)
}
fmt.Println(config.Mysql)
*/
var configFile []byte

type Config struct {
	DisablePathCorrection             bool   `yaml:"DisablePathCorrection"`
	EnablePathEscape                  bool   `yaml:"EnablePathEscape"`
	FireMethodNotAllowed              bool   `yaml:"FireMethodNotAllowed"`
	DisableBodyConsumptionOnUnmarshal bool   `yaml:"DisableBodyConsumptionOnUnmarshal"`
	TimeFormat                        string `yaml:"TimeFormat"`
	Charset                           string `yaml:"Charset"`
	Mysql                             Mysql  `yaml:"Mysql"`
}

type Mysql struct {
	Port   string `yaml:"Port"`
	Name   string `yaml:"Name"`
	Pwd    string `yaml:"Pwd"`
	Ip     string `yaml:"Ip"`
	DBName string `yaml:"DBName"`
}

func GetConfig() (e *Config, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("src/config/iris.yml")
	if err != nil {
		log.Fatalf("读取 yml 文件出错： %v ", err)
	}
}
