package config

type AppConfig struct {
    Profile string `yaml:"profile" env:"APP_PROFILE"`
    Name    string `yaml:"name" env:"APP_NAME" env-default:"GOUSS"`
    Port    int    `yaml:"port" env:"APP_PORT" env-default:"8080"`
    Host    string `yaml:"host" env:"APP_HOST" env-default:"0.0.0.0"`
}

type DatabaseConfig struct {
    Type     string `yaml:"type" env:"DB_TYPE" env-default:"in-memory"`
    Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
    Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
    Schema   string `yaml:"schema" env:"DB_SCHEMA" env-default:"postgres"`
    Username string `yaml:"user" env:"DB_USER" env-default:"user"`
    Password string `yaml:"pass" env:"DB_PASS" `
}

type Property struct {
	Env string `env:"ENV" env-default:"dev"`
	App AppConfig `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}