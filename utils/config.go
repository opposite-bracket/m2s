package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const fileName = "m2s.yaml"

type M2sConfig struct {
	Address         string `yaml:"address"`
	Mode            string `yaml:"mode"`
	RecorderLogPath string `yaml:"recorder_log_path"`
}

// GetConf will load m2s.yaml from the same
// path in which the binary was executed
func GetConf() (*M2sConfig, error) {
	dirPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configFile := fmt.Sprintf(
		"%s/%s",
		dirPath,
		"m2s.yaml",
	)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var data M2sConfig
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
