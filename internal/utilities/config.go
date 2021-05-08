package utilities

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "/.tjrc"
)

type Config struct {
	Context string `yaml:"context"`
}

func (c *Config) GetEmptyContext() *Config {
	return &Config{
		Context: "",
	}
}

func (c *Config) GetConfig() *Config {
	homeDir := os.Getenv("HOME")
	configPath := homeDir + configFile

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println(ColorString("Config file not exsists...", "orange", ""))

		emptyContext := c.GetEmptyContext()
		bytes, err := yaml.Marshal(emptyContext)
		if err != nil {
			log.Fatalf("error")
		}

		f, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Fatalf("Error", err)
		}

		f.Write(bytes)
		defer f.Close()

		fmt.Println(ColorString("Config file .tjrc created!", "green", ""))
		return emptyContext
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println("Cannot read config file", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshall config file error:", err)
	}

	return c
}

func (c *Config) UpdateConfig() {
	homeDir := os.Getenv("HOME")
	configPath := homeDir + configFile

	bytes, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error")
	}

	f, err := os.OpenFile(configPath, os.O_WRONLY, 0666)

	if err != nil {
		log.Println("Cannot update config file", err)
	}

	f.Write(bytes)
	defer f.Close()
}
