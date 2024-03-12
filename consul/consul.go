package consul

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/shenqingawei/framework/nacos"
	"net"
)

func RegisterConsul(port int, serviceName string) error {
	c, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("%v:%v", nacos.ConfigPz.Consul.Host, nacos.ConfigPz.Consul.Port),
	})
	if err != nil {
		return err
	}

	ip := GetIp()
	return c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    serviceName,      //todo:发现服务的名称
		Tags:    []string{"GRPC"}, //todo:连接服务类型
		Port:    port,             //todo:服务端口号
		Address: ip[0],            //todo:服务地址
		Check: &api.AgentServiceCheck{ //todo:健康检查
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", ip[0], port), //todo:健康检查地址
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
func GetIp() (ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range addrs {
		ipNet, isVailIpNet := addr.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = append(ip, ipNet.IP.String())
			}
		}

	}
	return ip
}
