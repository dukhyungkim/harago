package repository

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"harago/common"
	"harago/config"
	"log"
	"strings"
)

const sharedListKey = "/shared"

type Etcd struct {
	etcdClient *clientv3.Client
	sharedList map[string]struct{}
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

	ctx, cancel := context.WithTimeout(context.Background(), common.DefaultTimeout)
	defer cancel()

	resp, err := etcdClient.Get(ctx, sharedListKey)
	if err != nil {
		log.Println(fmt.Errorf("failed to get kv; %w", err))
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		log.Println(fmt.Errorf("failed to find value from key: %s", sharedListKey))
	}

	return &Etcd{etcdClient: etcdClient, sharedList: makeSharedList(string(resp.Kvs[0].Value))}, nil
}

func (e *Etcd) Close() {
	if err := e.etcdClient.Close(); err != nil {
		log.Println(err)
	}
}

func (e *Etcd) IsShared(name string) bool {
	_, has := e.sharedList[name]
	return has
}

func (e *Etcd) WatchSharedList() {
	watchChan := e.etcdClient.Watch(context.Background(), sharedListKey)

	for watchResp := range watchChan {
		if len(watchResp.Events) == 0 {
			continue
		}
		e.sharedList = makeSharedList(string(watchResp.Events[0].Kv.Value))
	}
}

func makeSharedList(s string) map[string]struct{} {
	fields := strings.Fields(s)
	sharedList := make(map[string]struct{})
	for _, s := range fields {
		sharedList[s] = struct{}{}
	}
	return sharedList
}
