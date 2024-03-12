package grpc

import (
	"github.com/shenqingawei/framework/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectionGRPC(serviceName string) (*grpc.ClientConn, error) {
	address, err := consul.FindConsulAddress(serviceName)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
