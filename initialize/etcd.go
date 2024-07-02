package initialize

import (
	"fmt"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/etcd"
	"time"
)

// EtcdClient Global variable to hold the etcd client
var EtcdClient *etcd.EtcdClient

// SetupEtcd initializes Etcd
func SetupEtcd() {
	etcdUrl := config.GetString("ETCD_URL", "http://localhost:2379")
	endpoints := []string{etcdUrl}
	dialTimeout := 5 * time.Second

	var err error
	// Create etcd client
	EtcdClient, err = etcd.NewEtcdClient(endpoints, dialTimeout)
	if err != nil {
		fmt.Println("Failed to create etcd client:", err)
	}
	defer func(EtcdClient *etcd.EtcdClient) {
		err := EtcdClient.Close()
		if err != nil {
			fmt.Println("Failed to close etcd client:", err)
		}
	}(EtcdClient)
}
