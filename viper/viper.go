package viper

import (
	"github.com/spf13/viper"
)

type NacosConf struct {
	Host  string
	Port  string
	Name  string
	Group string
}

var NacosConfig NacosConf

func InitViper() error {
	viper.SetConfigFile("/Users/dujiawei/go/src/week1/user/work/viper/viper.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	NacosConfig.Host = viper.GetString("NacosConfig.Host")
	NacosConfig.Port = viper.GetString("NacosConfig.Port")
	NacosConfig.Name = viper.GetString("NacosConfig.Name")
	NacosConfig.Group = viper.GetString("NacosConfig.Group")

	return nil
}
