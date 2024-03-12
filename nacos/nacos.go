package nacos

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	v "github.com/shenqingawei/framework/viper"
	"github.com/spf13/viper"

	"strconv"
)

func ConnectionNacos(path string) (config_client.IConfigClient, error) {
	err := v.InitViper(path)
	if err != nil {
		return nil, err
	}
	port := viper.GetString("NacosConfig.Port")
	host := viper.GetString("NacosConfig.Host")

	clientConfig := constant.ClientConfig{
		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: host,
			Port:   uint64(p),
		},
	}
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
}
func InitNaocs(path string) error {
	c, err := ConnectionNacos(path)
	if err != nil {
		return err
	}
	config, err := c.GetConfig(vo.ConfigParam{
		DataId: "user",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(config), &ConfigPz)
	if err != nil {
		return err
	}

	return nil
}

var ConfigPz C

type C struct {
	Mysql struct {
		DriverName           string `json:"DriverName"`
		Username             string `json:"Username"`
		Password             string `json:"Password"`
		ConnectionType       string `json:"ConnectionType"`
		Host                 string `json:"Host"`
		Port                 string `json:"Port"`
		DatabaseName         string `json:"DatabaseName"`
		ConnectionParameters string `json:"ConnectionParameters"`
	} `json:"Mysql"`
	Redis struct {
		Host string `json:"Host"`
		Port string `json:"Port"`
	} `json:"Redis"`
	Token struct {
		SecretKey string `json:"SecretKey"`
		Seconds   string `json:"Seconds"`
	} `json:"Token"`
	App struct {
		Host string `json:"Host"`
		Port string `json:"Port"`
	} `json:"app"`
	Consul struct {
		Host string `json:"Host"`
		Port string `json:"Port"`
	} `json:"consul"`
}
