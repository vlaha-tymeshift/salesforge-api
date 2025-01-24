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
	if err := sfCfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("salesforge-api config validation failed: %w", err)
	}
	return sfCfg, nil
}

func (c Config) Validate() error {
	if c.ServiceName == "" {
		return fmt.Errorf("service name is required")
	}
	if c.Environment == "" {
		return fmt.Errorf("environment is required")
	}
	if err := c.Server.Validate(); err != nil {
		return fmt.Errorf("server config validation failed: %w", err)
	}
	if err := c.MySql.Validate(); err != nil {
		return fmt.Errorf("mysql config validation failed: %w", err)
	}
	if err := c.Logger.Validate(); err != nil {
		return fmt.Errorf("logger config validation failed: %w", err)
	}
	return nil
}

func (c ServerConfig) Validate() error {
	if c.AppServerPort == 0 {
		return fmt.Errorf("app server port is required")
	}
	if c.HealthcheckPort == 0 {
		return fmt.Errorf("healthcheck port is required")
	}
	return nil
}

func (c MySqlConfig) Validate() error {
	if c.Db == "" {
		return fmt.Errorf("db name is required")
	}
	if c.User == "" {
		return fmt.Errorf("user is required")
	}
	if c.Pass == "" {
		return fmt.Errorf("password is required")
	}
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port == 0 {
		return fmt.Errorf("port is required")
	}
	return nil
}

func (c LoggerConfig) Validate() error {
	if c.Level == "" {
		return fmt.Errorf("level is required")
	}
	if c.Format == "" {
		return fmt.Errorf("format is required")
	}
	return nil
}
