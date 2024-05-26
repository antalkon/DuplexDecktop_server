package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DBConfig struct {
	DB `yaml:"db"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbName"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

func MustDatabase() *DBConfig {
	configPath := "config/dataBase.yaml"

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	var cfg DBConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("cannot unmarshal config: %s", err)
	}

	fmt.Println("Db final data:", &cfg)
	return &cfg
}
