package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const fileName = "m2s.yaml"

type m2sConfig struct {
	Mode string `yaml:"mode"`
	RecorderLogPath string `yaml:"recorder_log_path"`
}

func (c *m2sConfig) getConf() *m2sConfig {

	yamlFile, err := ioutil.ReadFile("fileName.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func NewConfig() {

}
