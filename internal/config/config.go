// Package config contains application configuration structures and data read logic.
//
// The package organizes configuration read and some platform-specific postprocessing.
//
// Main method is `ReadConfig(filepath)`. In general, it works the following way:
// 1. The method reads a file by the path `filepath` and tries to parse it as a YAML file into the structure `MainConfig`;
// 2. After that the method tries to load environment variables and overrides values from the YAML file for non-empty environment variables;
// 3. The package processes a platform-specific configuration actions, namely:
// 	- Heroku: if the environment variable `HEROKU` is set, the method overrides `MainConfig.Server.Port` value from `PORT` environment variable
package config

// MainConfig is a structure of the application configuration
// This describes a configuration file structure
// Each variable can be overridden with the environment variable
// To see the whole list of environment variables run application help (`wallet -h`)
type MainConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig is a set of application server configuration variables
// Each variable can be overridden with the environment variable
type ServerConfig struct {
	// Host is an application server host
	Host string `yaml:"host" env:"SERVER_HOST" env-description:"application server host"`
	// Host is an application server port
	Port string `yaml:"port" env:"SERVER_PORT" env-description:"application server port"`
}

// DatabaseConfig is a set of database configuration variables
// Each variable can be overridden with the environment variable
type DatabaseConfig struct {
	// DatabaseURL is an optional parameter that will contain a full connection option string.
	// Can be used on some cloud hostings like Heroku
	DatabaseURL *string `env:"DATABASE_URL" env-description:"database connection option string"`
	// Host of the database
	Host string `yaml:"host" env:"DATABASE_HOST" env-description:"database host"`
	// Port of the database
	Port string `yaml:"port" env:"PORT,DATABASE_PORT" env-description:"database port"`
	// Database name
	Database string `yaml:"database" env:"DATABASE_NAME" env-description:"database name"`
	// Username is a name of a database user
	Username string `yaml:"username" env:"DATABASE_USERNAME" env-description:"database user name"`
	// Password of a database user
	Password string `yaml:"password" env:"PASSWORD" env-description:"database user password"`
	SSL      string `yaml:"ssl" env:"DATABASE_SSL" env-description:"SSL status"`
	// ConnectionPool number of database connections in the application pool
	ConnectionPool int `yaml:"conn-pool" env:"DATABASE_CONN_POOL" env-description:"database connection pool size"`
	// ConnectionWait wait until the database will up in the infinite loop
	ConnectionWait bool `yaml:"conn-wait" env:"DATABASE_CONN_WAIT" env-description:"wait until database up"`
}
