package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl		string	`json:"db_url"`
	UserName	string	`json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.UserName = userName
	return write(*c)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filePath, jsonData, 0600)
}