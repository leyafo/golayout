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
func ServiceAdd(endpointsKey, addr string) error {
		em, err := endpoints.NewManager(etcdClientV3, endpointsKey)
		if err != nil{
			return err
		}
		return em.AddEndpoint(etcdClientV3.Ctx(), endpointsKey+"/"+addr, endpoints.Endpoint{Addr: addr});
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
func ServiceAddWithLease(lid clientv3.LeaseID, endpointsKey, addr string) error {
	em, err := endpoints.NewManager(etcdClientV3, endpointsKey)
	if err != nil{
		return err
	}
	return em.AddEndpoint(etcdClientV3.Ctx(), endpointsKey+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lid));
}

//ServiceDelete delete a service
func ServiceDelete(endpointsKey, addr string) error {
	em, err := endpoints.NewManager(etcdClientV3, endpointsKey)
	if err != nil{
		return err
	}
	return em.DeleteEndpoint(etcdClientV3.Ctx(), endpointsKey+"/"+addr)
}

//DialGrpc dial an RPC service using the etcd gRPC resolver and balancer:
func DialGrpc(endpointsKey, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	etcdResolver, err := resolver.NewBuilder(etcdClientV3)
	if err != nil{
		return nil, err
	}
	opts = append(opts, grpc.WithResolvers(etcdResolver))
	return  grpc.Dial("etcd:///" +endpointsKey + "/"+ serviceName, opts...)
}


