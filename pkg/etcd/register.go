package etcd


import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"

	//"google.golang.org/grpc"
)


var(
	etcdClientV3 *clientv3.Client
)

func InitETCD(url string)error{
	var err error
	etcdClientV3, err = clientv3.NewFromURL("http://localhost:2379")
	if err != nil{
		return err
	}

	return nil
}

func RegisterService(name, host string, port int)error{
	r := &etcdnaming.GRPCResolver{Client: etcdClientV3}
	addr := fmt.Sprintf("%s:%d", host, port)
	return r.Update(context.TODO(), name, naming.Update{Op: naming.Add, Addr: addr })
}


