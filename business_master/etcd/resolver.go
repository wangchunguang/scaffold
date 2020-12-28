package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/resolver"
)

const schema = "etcd"

// 实现grpc的grpc.resolve.Builder接口的Build与Scheme方法
type Resolver struct {
	endpoints []string
	service   string
	cli       *clientv3.Client
	cc        resolver.ClientConn
}

func NewResolver(endpoints []string, service string) resolver.Builder {
	return &Resolver{endpoints: endpoints, service: service}
}

// 返回etcd模式
func (r *Resolver) Scheme() string {
	//	最好用这种，因为grpc resolver.Register(r)在注册时，会取scheme，如果一个系统有多个grpc发现，就会覆盖之前注册的
	return schema + "_" + r.service
}

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

// Close
func (r *Resolver) Close() {
}

//  实现grpc.resolve.Builder接口的方法
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, ops resolver.BuildOptions) (resolver.Resolver, error) {
	var err error
	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints: r.endpoints,
	})
	if err != nil {
		log.Error("etcd resolver build err=", err)
		return nil, err
	}
	r.cc = cc
	go r.watch(fmt.Sprintf(r.service))
	return r, nil
}

// 监听
func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)
	update := func() {
		addrList := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrList = append(addrList, v)
		}
		r.cc.UpdateState(resolver.State{Addresses: addrList})
	}
	// clientv3.WithPrefix() 表示匹配key的前缀
	resp, err := r.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err == nil {
		for i, kv := range resp.Kvs {
			info := &ServiceInfo{}
			err := json.Unmarshal(kv.Value, info)
			if err != nil {
				log.Error("r.cli.Get error ")
			}
			addrDict[string(resp.Kvs[i].Value)] = resolver.Address{Addr: info.IP}
		}
	}
	update()
	// 获取clientv3.WithPrevKV() 上一个键值对
	watch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range watch {
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				info := &ServiceInfo{}
				err := json.Unmarshal(ev.Kv.Value, info)
				if err != nil {
					log.Error(err)
				} else {
					addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: info.IP}
				}
			case mvccpb.DELETE:
				delete(addrDict, string(ev.PrevKv.Key))
			}
		}
		update()
	}
}
