package main

type serviceConfig struct {
	APIConfig    apiConfig    `yaml:"api"`
	ConsulConfig consulConfig `yaml:"consul"`
}

type apiConfig struct {
	Port int `yaml:"port"`
}

type consulConfig struct {
	Host string `yaml:"host"`
}
