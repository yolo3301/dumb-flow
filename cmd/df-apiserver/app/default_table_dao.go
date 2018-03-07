package app

import (
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type DefaultTableDAO struct {
	ectdEndpoint string
}

func NewDefaultTableDAO() DefaultTableDAO {
	ep := os.Getenv("ETCD_ENDPOINT")
	if ep == "" {
		panic("Can't find ECTD_ENDPOINT in env vars")
	}

	return DefaultTableDAO{ectdEndpoint: ep}
}

func (dao *DefaultTableDAO) newEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{dao.ectdEndpoint},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	return cli, nil
}
