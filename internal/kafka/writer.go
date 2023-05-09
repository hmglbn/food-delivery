package kafka

import (
	"context"
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"log"
)

const shardKey = "1"

type Writer struct {
	kafkaWriter *kafka.Writer
}

func NewWriter(brokers []string, login, password, topic string) *Writer {
	mechanism, err := scram.Mechanism(scram.SHA256, login, password)
	if err != nil {
		log.Fatalln(err)
	}
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
		TLS: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
	return &Writer{
		kafkaWriter: &kafka.Writer{
			Addr:      kafka.TCP(brokers[0]),
			Topic:     topic,
			Balancer:  &kafka.Hash{},
			Transport: sharedTransport,
		},
	}
}

func (w *Writer) Write(msg []byte) error {
	if err := w.kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(shardKey),
		Value: msg,
	}); err != nil {
		return err
	}
	return nil
}
