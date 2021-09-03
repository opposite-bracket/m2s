package utils

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const fileName = "m2s.yaml"
const defaultPort = ":5000"

type M2sConfig struct {
	Address         string `yaml:"address"`
	Mode            string `yaml:"mode"`
	RecorderLogPath string `yaml:"recorder_log_path"`
}

// ValidateConfig will ensure the validity
// of config file
func ValidateConfig(config M2sConfig) error {

	if config.Address == "" {
		config.Address = defaultPort
	}

	if config.Mode != "record" && config.Mode != "mock" {
		return errors.New("invalid config mode: [valid: recorder|mock]")
	}

	if config.RecorderLogPath == "" {
		return errors.New("missing required config: [field: recorder_log_path]")
	}

	return nil
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

	var config M2sConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	if err := ValidateConfig(config); err != nil {
		return nil, err
	}

	return &config, nil
}
