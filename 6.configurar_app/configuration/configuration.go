package configuration

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Database struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
}

type Configuration struct {
	DB Database `yaml:"Database"`
}

func Load(filename string) (*Configuration, error) {

	config := &Configuration{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
