package initialize

import (
	"fmt"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/errors"
	"github.com/spf13/viper"
)

func Loadconfig() {

	viperconfig := viper.New()
	viperconfig.AutomaticEnv()
	viperconfig.SetDefault("MODE", "development")
	mode := viperconfig.GetString("MODE")
	fmt.Println("Mode::", mode)

	viperconfig.SetConfigType("yaml")
	viperconfig.SetConfigName(mode)
	viperconfig.AddConfigPath("./config")

	vipersecrets := viper.New()
	vipersecrets.SetConfigType("yaml")
	vipersecrets.SetConfigName(mode)
	vipersecrets.AddConfigPath("./secrets")

	config_err := viperconfig.ReadInConfig()
	if config_err != nil {
		errors.Must(global.Logger, config_err, "Error loading configuration")
	}

	secrets_err := vipersecrets.ReadInConfig()
	if secrets_err != nil {
		errors.Must(global.Logger, secrets_err, "Error loading configuration")
	}

	viperconfig.MergeConfigMap(vipersecrets.AllSettings())

	//read server configuration
	fmt.Println("Server Port::", viperconfig.GetInt("server.port"))

	// configure structure
	if err := viperconfig.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
