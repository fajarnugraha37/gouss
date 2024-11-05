package config

import (
	"fmt"
    "path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/fajarnugraha37/gouss/pkg/util"
)

var AppProperty Property

func init() {
	env := util.GetEnv("ENV", "dev")
	configLocation := filepath.Join(util.GetCwd(), fmt.Sprintf("config.%s.yaml", env))
	err := cleanenv.ReadConfig(configLocation, &AppProperty)
	util.FailOnError(err)
}