package config

import (
	"fmt"
	"reflect"
	"strings"
)

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

func (c *Property) Get(path string) (interface{}, error) {
	parts := strings.Split(path, ".")
	var value reflect.Value = reflect.
        ValueOf(c).
        Elem()

	for _, part := range parts {
		if value.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot access field %s in non-struct value", part)
        }

        value = value.FieldByName(part)
        if !value.IsValid() {
            return nil, fmt.Errorf("invalid field: %s", part)
        }
	}

	return value.Interface(), nil
}

func (c *Property) Set(path string, value interface{}) error {
    parts := strings.Split(path, ".")
    v := reflect.
        ValueOf(c).
        Elem()

    for i, part := range parts {
        if v.Kind() != reflect.Struct {
            return fmt.Errorf("invalid path: %s is not a struct", strings.Join(parts[:i], "."))
        }
        
        if i == len(parts)-1 {
            field := v.FieldByName(part)
            if !field.IsValid() {
                return fmt.Errorf("field not found: %s", part)
            }
            if !field.CanSet() {
                return fmt.Errorf("cannot set field: %s", part)
            }
            
            newValue := reflect.ValueOf(value)
            if !newValue.Type().AssignableTo(field.Type()) {
                return fmt.Errorf("cannot assign %v to %s (type %v)", value, part, field.Type())
            }
            field.Set(newValue)
         
            return nil
        }
        
        v = v.FieldByName(part)
        if !v.IsValid() {
            return fmt.Errorf("invalid field: %s", part)
        }
    }

    return fmt.Errorf("path does not lead to a field")
}
