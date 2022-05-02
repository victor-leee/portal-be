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

type Kubernetes struct {
	RegistryIP string
}

type Config struct {
	MysqlCfg *Mysql `yaml:"mysql_cfg"`
	K8SCfg   *Kubernetes
}

func Init() (*Config, error) {
	cfg := &Config{}
	file, err := os.Open("internal/config/cfg.yml")
	if err != nil {
		return nil, err
	}
	if err = yaml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}
	initFromEnv(cfg)

	return cfg, nil
}

func initFromEnv(config *Config) {
	if config.K8SCfg == nil {
		config.K8SCfg = &Kubernetes{}
	}
	if registryIP, ok := os.LookupEnv(EnvKubernetesRegistryIPKey); ok {
		config.K8SCfg.RegistryIP = registryIP
	}
}
