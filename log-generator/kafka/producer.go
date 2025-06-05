package kafka

import (
	"ErrorSense/log-generator/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type LogProducer interface {
	Produce(entry domain.LogEntry) error
	Close() error
}

type KProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(Broker string, topic string) *KProducer {
	return &KProducer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(Broker),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireOne,
			WriteTimeout: 10 * time.Millisecond,
		},
		topic: topic,
	}
}

func (p *KProducer) Produce(entry domain.LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("error al serializar LogEntry: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(entry.Source + ":" + entry.UserID),
		Value: data,
		Time:  entry.Timestamp,
	}

	err = p.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("error al escribir mensaje en Kafka: %w", err)
	}
	return nil
}

func (p *KProducer) Close() error {
	log.Println("Cerrando productor de Kafka...")
	return p.writer.Close()
}
