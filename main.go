package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	_, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("hello")
}
