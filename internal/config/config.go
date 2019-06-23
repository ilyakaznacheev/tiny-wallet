// Package config contains application configuration structures and data read logic
package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// MainConfig is a structure of the application configuration
// This describes a configuration file structure
// Each variable can be overridden with the environment variable
type MainConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig is a set of application server configuration variables
// Each variable can be overridden with the environment variable
type ServerConfig struct {
	// Host is an application server host
	Host string `yaml:"host" envconfig:"HOST" desc:"application server host"`
	// Host is an application server port
	Port string `yaml:"port" envconfig:"PORT" desc:"application server port"`
}

// DatabaseConfig is a set of database configuration variables
// Each variable can be overridden with the environment variable
type DatabaseConfig struct {
	// DatabaseURL is an optional parameter that will contain a full connection option string.
	// Can be used on some cloud hostings like Heroku
	DatabaseURL *string `envconfig:"URL" desc:"database connection option string"`
	// Host of the database
	Host string `yaml:"host" envconfig:"HOST" desc:"database host"`
	// Port of the database
	Port string `yaml:"port" envconfig:"PORT" desc:"database port"`
	// Database name
	Database string `yaml:"database" envconfig:"NAME" desc:"database name"`
	// Username is a name of a database user
	Username string `yaml:"username" envconfig:"USERNAME" desc:"database user name"`
	// Password of a database user
	Password string `yaml:"password" envconfig:"PASSWORD" desc:"database user password"`
	SSL      string `yaml:"ssl" envconfig:"SSL" desc:"SSL status"`
	// ConnectionPool number of database connections in the application pool
	ConnectionPool int `yaml:"conn-pool" envconfig:"CONN_POOL" desc:"database connection pool size"`
	// ConnectionWait wait until the database will up in the infinite loop
	ConnectionWait bool `yaml:"conn-wait" envconfig:"CONN_WAIT" desc:"wait until database up"`
}

// ReadConfig reads configuration from different sources
// 1. reads a configuration file and returns it's content in a structured manner
// 2. overriddes configuration with environment variables
func ReadConfig(path string) (*MainConfig, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg MainConfig

	// read YAML configuration file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	// read environment variables
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
