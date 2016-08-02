package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	log "github.com/Sirupsen/logrus"
)

type Metric string

type BackendConfig struct {
	DockerHost    string `yaml:"DockerHost"`
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

func parse(y []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(y, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func NewConfigFromByteArray(config []byte) (Config, error) {
	return parse(config)
}

func NewConfigFromFile(file string) (Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read config file %s: %v.", file, err)
		return Config{}, nil
	}
	return NewConfigFromByteArray(data)
}

func (c Config) Print() {
	fmt.Println("Backend: ", c.Backend)
	fmt.Println("BackendConfig")
	fmt.Println("\tDocker Host ", c.BackendConfig.DockerHost)
	fmt.Println("\tTLSCACert ", c.BackendConfig.TLSCACert)
	fmt.Println("\tTLSCert ", c.BackendConfig.TLSCert)
	fmt.Println("\tTLSKey ", c.BackendConfig.TLSKey)
	fmt.Println("\tAllowInsecure ", c.BackendConfig.AllowInsecure)
	fmt.Println("Metrics")

	for name, metric := range c.Metrics {
		fmt.Println("\t", name, " ", metric)
	}
}
