package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type EmailDeduperConfig struct {
	GenerateEmailCount  int     `yaml:"GENERATE_EMAIL_COUNT"`
	DuplicatePercentage float32 `yaml:"DUPLICATE_PERCENTAGE"`
}

// CreateFromFile populates a new Config with data from a YAML file
func CreateFromFile(config *EmailDeduperConfig, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, config)
}

// LoadConfigFromFile loads a EmailDeduperConfig from the file specified by path
func LoadConfigFromFile(path string) (EmailDeduperConfig, error) {
	config := EmailDeduperConfig{}
	err := CreateFromFile(&config, path)
	if err != nil {
		return config, err
	}
	return config, err
}

func LoadConfig() (EmailDeduperConfig, error) {
	return LoadConfigFromFile("config.yml")
}
