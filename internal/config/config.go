package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Log     `yaml:"log"`
	Server  `yaml:"server"`
	Storage `yaml:"storage"`
	Service `yaml:"service"`
}

// TODO add validator
func New(cfgPath string) (Config, error) {
	const f = "github.com/ulibaysya/krona/internal/daemon/config.New"

	file, err := os.Open(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", f, err)
	}
	defer func() {
		err = file.Close()
	}()

	// v := validator.New()
	// if err := v.RegisterValidation("global", validation); err != nil {
	// 	return Config{}, err
	// }

	cfg := Config{}
	// if err := yaml.NewDecoder(file, yaml.DisallowUnknownField(), yaml.Validator(v)).Decode(&cfg); err != nil {
	if err := yaml.NewDecoder(file, yaml.DisallowUnknownField()).Decode(&cfg); err != nil {
		// return Config{}, fmt.Errorf("%s: %s", f, yaml.FormatError(err, false, false))
		return Config{}, fmt.Errorf("%s: %w", f, err)
	}

	return cfg, err
}

// func validation(fl validator.FieldLevel) bool {
//
// 	return true
// }
