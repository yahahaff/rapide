package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// EtcdClient represents the etcd client wrapper.
type EtcdClient struct {
	client *clientv3.Client
}

// NewEtcdClient creates a new etcd client.
func NewEtcdClient(endpoints []string, dialTimeout time.Duration) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &EtcdClient{
		client: cli,
	}, nil
}

// Close closes the etcd client connection.
func (ec *EtcdClient) Close() error {
	return ec.client.Close()
}

// Put puts a key-value pair into etcd.
func (ec *EtcdClient) Put(ctx context.Context, key, value string) (*clientv3.PutResponse, error) {
	return ec.client.Put(ctx, key, value)
}

// Get gets the value for a given key from etcd.
func (ec *EtcdClient) Get(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return ec.client.Get(ctx, key)
}

// Delete deletes a key from etcd.
func (ec *EtcdClient) Delete(ctx context.Context, key string) (*clientv3.DeleteResponse, error) {
	return ec.client.Delete(ctx, key)
}
