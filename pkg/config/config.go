package config

import (
	"os"
	"strconv"

	"github.com/omnia-core/go-echo-template/pkg/log"

	"gopkg.in/yaml.v3"
)

const (
	localConfigPath = "config/config.yaml"
)

type Config struct {
	Listen     string    `yaml:"listen"`
	Env        string    `yaml:"env"`
	Postgresql DBConfig  `yaml:"postgresql"`
	JWTConfig  JWTConfig `yaml:"jwt"`
}

func (c Config) GetTestDBConfig(name string, port int) Config {
	return Config{
		Postgresql: DBConfig{
			Host:     "localhost",
			User:     name,
			Port:     port,
			Password: name,
			DBName:   name,
			Options: DBOptions{
				Connections: 100,
			},
		},
	}
}

func New() *Config {
	var config Config
	file, err := os.ReadFile(localConfigPath)
	if err != nil {
		log.New().Fatalf("err can not read local config file: %v", err)
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		log.New().Fatalf("err unmarshal yaml from local config file: %v", err)
	}

	return &config
}

type DBConfig struct {
	Host     string    `yaml:"host"`
	User     string    `yaml:"user"`
	Port     int       `yaml:"port"`
	Password string    `yaml:"password"`
	DBName   string    `yaml:"dbName"`
	Options  DBOptions `yaml:"options"`
}

func (d DBConfig) PortString() string {
	return strconv.Itoa(d.Port)
}

type DBOptions struct {
	Connections    int `yaml:"connections"`
	MinConnections int `yaml:"minConnections"`
	MaxConnections int `yaml:"maxConnections"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}
