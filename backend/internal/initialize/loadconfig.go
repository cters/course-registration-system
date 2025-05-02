package initialize

import (
	"fmt"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/errors"
	"github.com/spf13/viper"
)

func Loadconfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		errors.Must(global.Logger, err, "Error loading configuration")
	}

	//read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}