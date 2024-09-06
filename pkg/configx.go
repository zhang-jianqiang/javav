package pkg

import (
	"os"
	"path/filepath"
)

var configName = "javav.txt"

func SetConfig(path string) (err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configFile := filepath.Join(home, configName)
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(path)
	if err != nil {
		return
	}

	return nil
}

func ReadConfig() (path string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configFile := filepath.Join(home, configName)
	readByte, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	return string(readByte), nil
}
