package etcd

import (
	"context"
	"golayout/pkg/logger"
	"google.golang.org/grpc"
	"time"
)

var (
	grpcConnections = make(map[string]*grpc.ClientConn)
)

func GetGrpcConnection(endpointsKey, serviceName string)(conn *grpc.ClientConn, err error){
	var connExisted bool
	conn, connExisted = grpcConnections[serviceName]
	if connExisted && conn != nil{
		return conn, nil
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	logger.Infof("dial %s/%s...", endpointsKey, serviceName)
	conn, err = DialGrpc(timeoutCtx, endpointsKey, serviceName, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		return nil, err
	}

	grpcConnections[serviceName] = conn
	return conn, err
}