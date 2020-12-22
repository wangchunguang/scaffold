package etcd

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// 服务信息
type ServiceInfo struct {
	Name string // etcd名称
	IP   string // etcd ip
}

type Service struct {
	ServiceInfo ServiceInfo
	stop        chan error
	leaseId     clientv3.LeaseID // 租约编号
	client      *clientv3.Client // 客户端提供和管理etcd v3客户端会话。
}

//  NewService 创建一个注册服务
func NewService(info ServiceInfo, endpoints []string) (service *Service, err error) {
	client, err := clientv3.New(clientv3.Config{ // 根据配置创建一个新的 etcd3客户端
		Endpoints:   endpoints, // 端点是URL列表。
		DialTimeout: time.Second * 200,
	})
	if err != nil {
		log.Error("etcd client err", err)
		return nil, err
	}
	service = &Service{
		ServiceInfo: info,
		client:      client,
	}
	return
}

// Start 注册服务启动
func (service *Service) Start() (err error) {
	ch, err := service.keepAlive()
	if err != nil {
		log.Error("etcd start err=", err)
		return err
	}
	for {
		select {
		case err := <-service.stop:
			return err
		case <-service.client.Ctx().Done():
			return errors.New("service closed")
		case resp, ok := <-ch:
			//	 监听租约
			if !ok {
				log.Info("keep alive channel closed")
				return service.revoke()
			}
			log.Info("recv reply from service key=", service.getKey(), "ttl =", resp.TTL)
		}
	}
	return
}

// 停止
func (service *Service) Stop() {
	service.stop <- nil
}

// 续约
func (service *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	info := &service.ServiceInfo
	key := info.Name + "/" + info.IP
	val, _ := json.Marshal(info)

	// 创建一个租约
	resp, err := service.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Error("etcd grant err:", err)
		return nil, err
	}
	// 新增一个kv键值对
	_, err = service.client.Put(context.TODO(), key, string(val), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Error("etcd put kv err=", err)
		return nil, err
	}
	service.leaseId = resp.ID
	return service.client.KeepAlive(context.TODO(), resp.ID)
}

// 撤销
func (service *Service) revoke() error {
	_, err := service.client.Revoke(context.TODO(), service.leaseId)
	if err != nil {
		log.Error("etcd revoke err=", err)
		return err
	}
	return nil
}

// 获取key
func (service *Service) getKey() string {
	return service.ServiceInfo.Name + "/" + service.ServiceInfo.IP
}
