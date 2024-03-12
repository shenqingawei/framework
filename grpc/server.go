package grpc

import (
	"fmt"
	"github.com/shenqingawei/framework/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RegisterGRPC(port int, address, serviceName string, tagsType []string, fuc func(r *grpc.Server)) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}
	err = consul.RegisterConsul(port, address, serviceName, tagsType) //todo:服务注册
	if err != nil {
		return err
	}
	s := grpc.NewServer()                                      //todo:GRPC 实例
	grpc_health_v1.RegisterHealthServer(s, health.NewServer()) //todo:健康检查
	reflection.Register(s)                                     //todo:反射端口号到本地服务
	fuc(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
