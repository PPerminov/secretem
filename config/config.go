package config

import (
	"chat/logAndErrors"
	yaml "gopkg.in/yaml.v2"
	"os"
)

type Configuration struct {
	PublicKeyPath  string `yaml:"public_key_path"`
	PrivateKeyPath string `yaml:"private_key_path"`
	KeyType        string `yaml:"key_type"`
}

func LoadConfigFile(file string) *Configuration {
	fileBytes, err := os.ReadFile(file)
	logAndErrors.IfErrorThenFatal(err, "")
	config := Configuration{}
	err = yaml.Unmarshal(fileBytes, &config)
	logAndErrors.IfErrorThenFatal(err, "")
	return &config
}
