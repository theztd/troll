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
	Game        Game `yaml:"game"`
	//	Endpoints   []Endpoint
	Endpoints []Endpoint
}

type Game struct {
	Route        string   `yaml:"route"`
	TemplatePath string   `yaml:"templatePath"`
	Backends     []string `yaml:"backends"`
}

type Endpoint struct {
	Kind     string
	Method   string
	Path     string
	Query    string
	Args     []string
	Code     int
	Response string
	DSN      string `yaml:"dsn"`
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

	if LOG_LEVEL == "debug" {
		log.Printf("DEBUG [Loaded config]:\n%v+\n", yamlData)
	}

	return yamlData
}
