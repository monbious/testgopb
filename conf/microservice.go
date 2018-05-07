package conf

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type ConfigServiceMicro struct {
	ServiceAddress string                `yaml:"serviceAddress"`
	ConsulAddress  string                `yaml:"consulAddress"`
	ServiceName    string                `yaml:"serviceName"`
}

func LoadConf(conf interface{}, filePath string) {
	absPath, err := filepath.Abs(filePath)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println(err)
	}
	e := yaml.Unmarshal([]byte(data), conf)
	if e != nil {
		fmt.Println("配置文件解析失败")
		return
	}
}