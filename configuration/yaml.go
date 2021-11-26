package configuration

import (
	"cx-20-api/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func GetEnv() entity.Config {
	filename, _ := filepath.Abs("env-prod.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config entity.Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
