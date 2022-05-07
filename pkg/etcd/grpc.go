package etcd

import (
	"context"
	"errors"
	"bridgeswap/pkg/logger"
	"google.golang.org/grpc"
	"sync"
	"sync/atomic"
)

type policyType int

const (
	RoundRobin policyType = iota + 1
)

type LoadBalancer struct {
	policy  policyType
	servers []string

	key  string
	ch   chan struct{}
	done chan struct{}

	count uint64
	lock  sync.Mutex
}

func NewLoadBalancer(key string, policy policyType) (*LoadBalancer, error) {
	if etcdClientV3 == nil {
		return nil, errors.New("please call Init first")
	}
	logger.Info("watch key is: ", key)
	lb := &LoadBalancer{policy: policy, key: key}
	var err error
	lb.servers, err = etcdClientV3.GetEntries(lb.key)
	if err != nil {
		panic(err)
	}

	lb.done = make(chan struct{})
	lb.ch = make(chan struct{})

	go etcdClientV3.WatchPrefix(key, lb.ch)
	go lb.update()

	return lb, nil
}

func (lb *LoadBalancer) GetEntry() string {
	if len(lb.servers) == 0 {
		return ""
	}
	if lb.policy == RoundRobin {
		count := atomic.AddUint64(&lb.count, 1)
		return lb.servers[int(count)%len(lb.servers)]
	}

	return lb.servers[0]
}

func (lb *LoadBalancer) update() {
	logger.Info(lb.servers)
	for {
		select {
		case <-lb.ch:
			entries, err := etcdClientV3.GetEntries(lb.key)
			if err != nil {
				logger.Error("call etcd get entries failed: ", err)
			}
			lb.lock.Lock()
			lb.servers = entries
			lb.lock.Unlock()
		case <-lb.done:
			break
		}
	}
}

func (lb *LoadBalancer) Close() {
	close(lb.done)
	close(lb.ch)
}

func DialGrpc(ctx context.Context, lb *LoadBalancer, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	s := lb.GetEntry()
	if s == "" {
		return nil, errors.New("load balance get an empty string")
	}

	logger.Infof("dial %s ...", s)
	return grpc.DialContext(ctx, s, opts...)
}
