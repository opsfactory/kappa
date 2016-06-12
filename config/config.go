package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Metric string

type BackendConfig struct {
	TLSCACert     string `yaml:"TLSCACert"`
	TLSCert       string `yaml:"TLSCert"`
	TLSKey        string `yaml:"TLSKey"`
	AllowInsecure bool   `yaml:"AllowInsecure"`
}

type Config struct {
	Backend       string            `yaml:"backend"`
	BackendConfig BackendConfig     `yaml:"backend_config"`
	Metrics       map[string]Metric `yaml:"metrics"`
}

func Parse(y []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(y, &c)

	if err != nil {
		return c, err
	}

	if err != nil {
		return c, err
	}

	return c, nil
}

func Print(c Config) {
	fmt.Println("Backend: ", c.Backend)
	fmt.Println("BackendConfig")
	fmt.Println("\tTLSCACert ", c.BackendConfig.TLSCACert)
	fmt.Println("\tTLSCert ", c.BackendConfig.TLSCert)
	fmt.Println("\tTLSKey ", c.BackendConfig.TLSKey)
	fmt.Println("\tAllowInsecure ", c.BackendConfig.AllowInsecure)
	fmt.Println("Metrics")

	for name, metric := range c.Metrics {
		fmt.Println("\t", name, " ", metric)
	}
}
