package etcd

import (
	"context"
	"golayout/pkg/logger"
	"google.golang.org/grpc"
	"sync"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
)

var (
	etcdClientV3 etcdv3.Client = nil
	once         sync.Once
)

//Init init etcd service
func Init(endpoints []string) error {
	var err error
	once.Do(func() {
		etcdClientV3, err = etcdv3.NewClient(context.TODO(), endpoints,
			etcdv3.ClientOptions{DialOptions: []grpc.DialOption{grpc.WithBlock()}})
	})
	return err
}

//ServiceAdd add a endpoint service
func ServiceAdd(key, value string) error {
	logger.Infof("add service key=%s value=%s", key, value)
	return etcdClientV3.Register(etcdv3.Service{
		Key:   key,
		Value: value,
		TTL:   etcdv3.NewTTLOption(3*time.Second, 10*time.Second),
	})
}

func ServiceDelete(key string) error {
	return etcdClientV3.Deregister(etcdv3.Service{Key: key})
}

//ServiceList list registered services
func ServiceList(prefix string) ([]string, error) {
	return etcdClientV3.GetEntries(prefix)
}
