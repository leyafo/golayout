package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"time"
)

var(
	etcdClientV3 *clientv3.Client
)

//InitEtcd init etcd service
func InitEtcd(etcdConfig clientv3.Config)error{
	if etcdConfig.DialTimeout == 0 {
		etcdConfig.DialTimeout = 3 * time.Second
	}
	if etcdConfig.DialKeepAliveTime == 0 {
		etcdConfig.DialKeepAliveTime = 3 * time.Second
	}
	if etcdConfig.Context == nil{
		etcdConfig.Context = context.TODO()
	}

	var err error
	etcdClientV3, err = clientv3.New(etcdConfig)
	return err
}

//ServiceAdd add a endpoint service
func ServiceAdd(endpointKey, addr string) error {
		em, err := endpoints.NewManager(etcdClientV3, endpointKey)
		if err != nil{
			return err
		}
		return em.AddEndpoint(etcdClientV3.Ctx(), endpointKey+"/"+addr, endpoints.Endpoint{Addr: addr});
}

//ServiceList list registered services
func ServiceList(endpointsKey string)([]string, error){
	em, err := endpoints.NewManager(etcdClientV3, endpointsKey)
	if err != nil{
		return nil, err
	}

	keyMap, err := em.List(etcdClientV3.Ctx())
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range(keyMap){
		result = append(result, v.Addr)
	}

	return result, err
}


//ServiceAddWithLease add a endpoint service with lease
func ServiceAddWithLease(c *clientv3.Client, lid clientv3.LeaseID, service, addr string) error {
	em, err := endpoints.NewManager(c, service)
	if err != nil{
		return err
	}
	return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr:addr}, clientv3.WithLease(lid));
}

//ServiceDelete delete a service
func ServiceDelete(service, addr string) error {
	em, err := endpoints.NewManager(etcdClientV3, service)
	if err != nil{
		return err
	}
	return em.DeleteEndpoint(etcdClientV3.Ctx(), service+"/"+addr)
}

//DialGrpc dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
func DialGrpc(endpointKey string) (*grpc.ClientConn, error) {
	etcdResolver, err := resolver.NewBuilder(etcdClientV3);
	if err != nil{
		return nil, err
	}
	return  grpc.Dial("etcd:///" + endpointKey, grpc.WithResolvers(etcdResolver))
}


