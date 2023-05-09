package kafka

import (
	"context"
	"crypto/tls"
	"food-delivery/internal/helper"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func NewReader(ctx context.Context, brokers []string, login, password, groupID, topic string) *kafka.Reader {
	mechanism, err := scram.Mechanism(scram.SHA256, login, password)
	if err != nil {
		log.Fatalln(err)
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
		Dialer: &kafka.Dialer{
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		},
	})

	if r != nil {
		go func() {
			select {
			case <-ctx.Done():
				helper.Log("kafka context done")
			}
		}()
	}

	return r
}
