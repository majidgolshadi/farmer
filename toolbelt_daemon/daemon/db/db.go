package db

import (
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
	"time"
)


var kapi client.KeysAPI


func Connect() error {
	c, err := client.New(client.Config{
		Endpoints:               []string{"http://etcd:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})

	if err != nil {
		return err
	}

	kapi = client.NewKeysAPI(c)

	return nil
}

func Set(key string, value string) error {
	_, err := kapi.Set(context.Background(), key, value, nil)
	return err
}

func Get(key string) (string, error) {
	response, err := kapi.Get(context.Background(), key, nil)
	if  err != nil {
		return "", err
	}

	return response.Node.Value, nil
}
