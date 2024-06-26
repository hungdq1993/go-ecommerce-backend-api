package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Using config file:", viper.GetInt("server.port"))
	fmt.Println("Using config file:", viper.GetString("server.host"))

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling confi  %v", err)
	}

	fmt.Println("Using config file:", config.Server.Port)

	for _, db := range config.Database {
		fmt.Println("Using config file:", db.Host)
		fmt.Println("Using config file:", db.Port)
		fmt.Println("Using config file:", db.User)

	}
}
