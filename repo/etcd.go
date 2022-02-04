package repo

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"harago/common"
	"harago/config"
	"log"
)

type Etcd struct {
	etcdClient *clientv3.Client
}

func NewEtcd(cfg *config.Etcd) (*Etcd, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: common.DefaultTimeout,
		Username:    cfg.Username,
		Password:    cfg.Password,
	})
	if err != nil {
		return nil, common.ErrConnEtcd(err)
	}

	return &Etcd{etcdClient: etcdClient}, nil
}

func (e *Etcd) Close() {
	if err := e.etcdClient.Close(); err != nil {
		log.Println(err)
	}
}
