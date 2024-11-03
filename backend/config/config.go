package config

import (
	"fmt"
    "path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/fajarnugraha37/gouss/pkg/util"
)

var AppProperty Property

func init() {
	err := cleanenv.ReadEnv(&AppProperty)
	util.FailOnError(err)
	
	configLocation := filepath.Join(util.GetCwd(), fmt.Sprintf("config.%s.yaml", AppProperty.Env))
	err = cleanenv.ReadConfig(configLocation, &AppProperty)
	util.FailOnError(err)
}