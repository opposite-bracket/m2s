package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const PERMISSION = 0644
const PREFIX = ""
const INDENT = "  "

const OutputDir = "/tmp/m2s"

func CreateOutDir() error {
	return os.MkdirAll(OutputDir, os.ModePerm)
}

func WriteToJsonFile(fileRelPath string, data interface{}) error {
	file, err := json.MarshalIndent(data, PREFIX, INDENT)
	if err != nil {
		log.Printf("failed to encode data to json: [error: %s]", err)
		return err
	}

	filePath := fmt.Sprintf(
		"%s/%s",
		OutputDir,
		fileRelPath,
	)

	err = ioutil.WriteFile(filePath, file, PERMISSION)

	if err != nil {
		return err
	}

	return nil
}
