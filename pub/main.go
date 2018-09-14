package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
	"time"
	"fmt"
)

const (
	clientID  = "test-publisher"
	clusterID = "test-cluster"
	subject   = "test_subject"
	natsURL   = "nats://natss:4222"
)

func main() {
	<-time.After(20 * time.Second)

	ticker := time.NewTicker(time.Second * 10)

	for range ticker.C {
		m := time.Now().Unix()
		if err := publish(fmt.Sprintf("%d", m)); err != nil {
			log.Fatal(err)
		}
		log.Printf("published %d", m)
	}
}

func publish(msg string) error {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Publish(subject, []byte(msg))
}