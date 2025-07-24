package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlRoot struct {
	Name        string
	Description string
	Version     string
	//	Endpoints   []Endpoint
	Endpoints []Endpoint
}

type Endpoint struct {
	Kind     string
	Method   string
	Path     string
	Query    string
	Args     []string
	Code     int
	Response string
}

func LoadYaml(path string) (yamlData YamlRoot) {
	yamlF, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	err2 := yaml.Unmarshal(yamlF, &yamlData)
	if err2 != nil {
		log.Fatalln(err2)
	}

	return yamlData
}
