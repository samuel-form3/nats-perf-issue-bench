package main

import (
	"fmt"
	"sync/atomic"

	"github.com/form3tech-oss/f1/pkg/f1"
	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/nats-io/nats.go"
)

var sent int32

func main() {
	f1Scenarios := f1.Scenarios().Add("jetstream", jetstreamTest)
	f1Scenarios.Execute()
}

func jetstreamTest(t *testing.T) testing.RunFn {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Create the stream
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     "test",
		Subjects: []string{"test.*.subj"},
		Replicas: 3,
	})

	if err != nil {
		t.Error(err)
	}

	// Purge the stream
	js.PurgeStream("test")

	// Create consumers
	_, err = js.AddConsumer("test", &nats.ConsumerConfig{
		AckPolicy: nats.AckExplicitPolicy,
		Durable:   "test",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	sub, err := js.Subscribe("test.*.subj", HandleMessages, nats.DeliverNew(), nats.ManualAck())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Cleanup(func() {
		sub.Drain()
		sub.Unsubscribe()
		nc.Close()
	})

	msg := make([]byte, 1024)
	runFn := func(t *testing.T) {
		msgID := atomic.AddInt32(&sent, 1)
		_, err := js.Publish(fmt.Sprintf("test.%d.subj", msgID), msg)
		if err != nil {
			t.Error(err)
		}
	}
	return runFn
}

func HandleMessages(m *nats.Msg) {
	if err := m.Ack(); err != nil {
		fmt.Println("error acking", err)
	}
}
