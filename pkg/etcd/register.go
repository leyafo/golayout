package etcd

import (
	"context"
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var(
	etcdClientV3 *clientv3.Client
	cancelFunc context.CancelFunc

	EmptyKey = errors.New("Empty key")
)

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


// GetEntries implements the etcd Client interface.
func GetEntries(key string) ([]string, error) {
	if key == ""{
		return nil, EmptyKey
	}
	resp, err := etcdClientV3.KV.Get(etcdClientV3.Ctx(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	entries := make([]string, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		entries[i] = string(kv.Value)
	}

	return entries, nil
}

// WatchPrefix Watch prefix
func WatchPrefix(prefix string, ch chan struct{}) {
	var ctx context.Context
	ctx, cancelFunc  = context.WithCancel(etcdClientV3.Ctx())
	watcher := clientv3.NewWatcher(etcdClientV3)

	wch := watcher.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithRev(0))
	ch <- struct{}{}
	for wr := range wch {
		if wr.Canceled {
			return
		}
		ch <- struct{}{}
	}
}

func Register(key, value string) error {
	var err error

	leaser := clientv3.NewLease(etcdClientV3)

	grantResp, err := leaser.Grant(etcdClientV3.Ctx(), int64(time.Second*30))
	if err != nil {
		return err
	}

	_, err =  etcdClientV3.Put(etcdClientV3.Ctx(), key, value, clientv3.WithLease(grantResp.ID))

	if err != nil {
		return err
	}

	// this will keep the key alive 'forever' or until we revoke it or
	// the context is canceled
	hbch, err := leaser.KeepAlive(etcdClientV3.Ctx(), grantResp.ID)
	if err != nil {
		return err
	}

	// discard the keepalive response, make etcd library not to complain
	// fix bug #799
	go func() {
		for {
			select {
			case r := <-hbch:
				// avoid dead loop when channel was closed
				if r == nil {
					return
				}
			case <-etcdClientV3.Ctx().Done():
				return
			}
		}
	}()

	return nil
}

func Deregister(key string) error {
	if key == ""{
		return EmptyKey
	}
	if _, err := etcdClientV3.Delete(etcdClientV3.Ctx(), key, clientv3.WithIgnoreLease()); err != nil {
		return err
	}

	return nil
}