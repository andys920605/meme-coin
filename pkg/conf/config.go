package conf

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/andys920605/meme-coin/pkg/errors"
)

type Config struct {
	Server struct {
		Name    string `validate:"required"`
		Port    string `validate:"required"`
		Version string `validate:"required"`
	}
	Log struct {
		Level string `validate:"required"`
	}
	MySQL struct {
		Host     string `validate:"required"`
		Port     int    `validate:"required"`
		Username string `validate:"required"`
		Password string `validate:"required"`
		Database string `validate:"required"`
		MaxIdle  int    `validate:"required"`
		MaxOpen  int    `validate:"required"`
	}
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config error")
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "marshal error")
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, errors.Wrap(err, "validate error")
	}

	return &config, nil
}
