package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	MysqlCfg *Mysql `yaml:"mysql_cfg"`
}

func Init() (*Config, error) {
	cfg := &Config{}
	file, err := os.Open("cfg.yml")
	if err != nil {
		return nil, err
	}
	if err = yaml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
