package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	setting "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/pkg"
)

func LoadConfig() (config setting.Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()	
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, fmt.Errorf("config file not found: %w", err)
		}
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode config into struct: %w", err)
	}
	return
}