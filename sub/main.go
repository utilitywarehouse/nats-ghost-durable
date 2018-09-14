package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
	"os"
	"encoding/hex"
	"crypto/rand"
	"os/signal"
	"syscall"
	"context"
	"github.com/pkg/errors"
	"regexp"
	"time"
	"github.com/nats-io/go-nats-streaming/pb"
)

const (
	clientID    = "test-publisher"
	clusterID   = "test-cluster"
	subject     = "test_subject"
	natsURL     = "nats://natss:4222"
)

var nameRx = regexp.MustCompile("/([^/]+)/master")

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal)
		signal.Notify(sCh, os.Interrupt, syscall.SIGTERM)
		<-sCh
		log.Println("shutting down")
		cancel()
	}()

	<-time.After(5 * time.Second)

	res := nameRx.FindStringSubmatch(os.Getenv("MASTER_NAME"))
	name := res[1]

	errCh := make(chan error)

	connectionLostHandler := func(_ stan.Conn, e error) {
		errCh <- errors.Wrap(e, "nats streaming connection lost")
	}

	conn, err := stan.Connect(clusterID, clientID+generateID(), stan.NatsURL(natsURL), stan.SetConnectionLostHandler(connectionLostHandler))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	handler := func(m *stan.Msg) {
		log.Printf("received %s (seq = %d)", string(m.Data), m.Sequence)
		m.Ack()
	}

	sub, err := conn.QueueSubscribe(
		subject,
		name,
		handler,
		stan.DurableName(name),
		stan.SetManualAckMode(),
		stan.MaxInflight(1),
		stan.StartAt(pb.StartPosition_First),
	)
	if err != nil {
		log.Panic(err)
	}
	defer sub.Close()

	log.Printf("subscribed with durable name %s..", name)

	select {
	case <-ctx.Done():
		log.Println("graceful shutdown")
	case err := <-errCh:
		log.Printf("error: %s", err)
	}
}

func generateID() string {
	random := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	_, err := rand.Read(random)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(random)
}
