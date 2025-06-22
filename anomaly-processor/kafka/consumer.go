package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"ErrorSense/log-generator/domain"
	"github.com/segmentio/kafka-go"
)

type LogConsumer interface {
	Start(ctx context.Context, handler func(entry domain.LogEntry) error) error
	Close() error
}

type KConsumer struct {
	reader  *kafka.Reader
	topic   string
	groupID string
}

func NewKafkaConsumer(broker, topic, groupID string) *KConsumer {
	return &KConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:          []string{broker},
			Topic:            topic,
			GroupID:          groupID,
			MinBytes:         10e3,
			MaxBytes:         10e6,
			MaxAttempts:      3,
			ReadBatchTimeout: 5 * time.Second,
			Logger:           kafka.LoggerFunc(log.Printf),
			ErrorLogger:      kafka.LoggerFunc(log.Printf),
		}),
		topic:   topic,
		groupID: groupID,
	}
}

func (c *KConsumer) Start(ctx context.Context, handler func(entry domain.LogEntry) error) error {
	log.Printf("Iniciando el consumidor de Kafka para el tópico: %s, grupo: %s", c.topic, c.groupID)
	for {
		select {
		case <-ctx.Done():
			log.Println("Consumidor detenido por cancelación de contexto.")
			return nil
		default:

		}

		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			// Diferenciar entre cancelación de contexto y otros errores de lectura
			if ctx.Err() != nil {
				return nil
			}
			log.Printf("Error al leer mensaje de Kafka: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		var entry domain.LogEntry
		if err := json.Unmarshal(m.Value, &entry); err != nil {
			log.Printf("Error al deserializar mensaje JSON del tópico '%s' (offset %d): %v. Mensaje: %s",
				c.topic, m.Offset, err, string(m.Value))

			if err := c.reader.CommitMessages(ctx, m); err != nil {
				log.Printf("Error al hacer commit de mensaje corrupto en Kafka: %v", err)
			}
			continue
		}

		// Pasa el LogEntry deserializado al handler de la lógica de negocio
		if err := handler(entry); err != nil {
			log.Printf("Error al procesar log entry de la lógica de negocio: %v. Mensaje: %+v", err, entry)
		}

		// Confirma que el mensaje ha sido procesado (offset actualizado)
		if err := c.reader.CommitMessages(ctx, m); err != nil {
			log.Printf("Error al hacer commit de mensaje procesado en Kafka: %v", err)
		}
	}
}

func (c *KConsumer) Close() error {
	log.Println("Cerrando consumidor de Kafka...")
	return c.reader.Close()
}
