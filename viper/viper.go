package viper

import (
	"github.com/spf13/viper"
)

func InitViper(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
