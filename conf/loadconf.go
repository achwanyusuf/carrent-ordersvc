package conf

import (
	"encoding/json"

	"github.com/spf13/viper"
)

func New(file string) (Config, error) {
	var cfg Config
	viper.AddConfigPath(".")
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	a, _ := json.Marshal(cfg)

	return cfg, nil
}
