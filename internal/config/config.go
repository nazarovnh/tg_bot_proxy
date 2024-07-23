package config

import (
	"os"
	"tg_bot_proxy/internal/database"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Bot struct {
	Token       string `yaml:"token"`
	WebhookLink string `yaml:"webhookLink"`
}

type Config struct {
	Server    Server                   `yaml:"server"`
	Bot       Bot                      `yaml:"bot"`
	Database  database.Config          `yaml:"database"`
	// Migration database.MigrationConfig `yaml:"migration"`
}


func GetConfig(configPath string, cfg interface{}) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return err
	}

	return nil
}