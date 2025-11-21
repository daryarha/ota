package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Host   map[string]string `json:"host"`
	Server Server            `json:"server"`
}

type Server struct {
	Host string `json:"host"`
}

func GetConfiguration() (conf Configuration, err error) {
	c, err := os.ReadFile("data/config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(c, &conf)
	if err != nil {
		return
	}
	return
}

func (c Configuration) GetHost(provider string) string {
	if host, ok := c.Host[provider]; ok {
		return host
	}
	return ""
}
