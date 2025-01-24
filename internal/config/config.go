package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
)

type Config struct {
	ServiceName string       `yaml:"ServiceName"`
	Environment string       `yaml:"Environment"`
	Server      ServerConfig `yaml:"Server"`
	MySql       MySqlConfig  `yaml:"MySql"`
	Logger      LoggerConfig `yaml:"Logger"`
}

type ServerConfig struct {
	AppServerPort     int  `yaml:"AppServerPort"`
	HealthcheckPort   int  `yaml:"HealthcheckPort"`
	JWTAuthentication bool `yaml:"JWTAuthentication"`
}

type MySqlConfig struct {
	Db   string `yaml:"Db"`
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type LoggerConfig struct {
	Level  string `yaml:"Level"`
	Format string `yaml:"Format"`
}

func LoadFromFilesystem(filesystem fs.FS, path string) (cfg Config, err error) {
	f, err := filesystem.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to open config file: %w", err)
	}
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("decoding yaml failed: %w", err)
	}
	return
}

func LoadConfig(filesystem fs.FS) (cfg Config, err error) {
	sfCfg, err := LoadFromFilesystem(filesystem, "config/config.yaml")
	if err != nil {
		return Config{}, fmt.Errorf("fail to load salesforge-api config: %w", err)
	}

	return sfCfg, nil
}
