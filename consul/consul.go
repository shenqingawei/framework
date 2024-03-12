package consul

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/shenqingawei/framework/nacos"
)

func RegisterConsul(port int, address, serviceName string, tagsTYpe []string) error {
	c, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", nacos.ConfigPz.Consul.Host, nacos.ConfigPz.Consul.Port),
	})
	if err != nil {
		return err
	}
	return c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    serviceName, //todo:发现服务的名称
		Tags:    tagsTYpe,    //todo:连接服务类型
		Port:    port,        //todo:服务端口号
		Address: address,     //todo:服务地址
		Check: &api.AgentServiceCheck{ //todo:健康检查
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", address, port), //todo:健康检查地址
			DeregisterCriticalServiceAfter: "30s",
		},
	})
}
func FindConsulAddress(serviceName string) (string, error) {
	c, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", nacos.ConfigPz.Consul.Host, nacos.ConfigPz.Consul.Port),
	})
	if err != nil {
		return "", err
	}
	name, data, err := c.Agent().AgentHealthServiceByName(serviceName) //todo:发现服务地址
	if err != nil {
		return "", err
	}
	if name != "passing" {
		return "", errors.New("服务健康检查未通过")
	}
	return fmt.Sprintf("%v:%v", data[0].Service.Address, data[0].Service.Port), nil
}
