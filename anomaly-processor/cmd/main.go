package cmd

import (
	"ErrorSense/anomaly-processor/kafka"
	"ErrorSense/log-generator/domain"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error al cargar el archivo .env: %v", err)
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	if kafkaBroker == "" || kafkaTopic == "" || kafkaGroupID == "" {
		log.Fatalf("Las variables de entorno KAFKA_BROKER, KAFKA_TOPIC o KAFKA_GROUP_ID no están configuradas.")
	}

	log.Printf("Iniciando Anomaly-Processor con Broker: %s, Topic: %s, Group ID: %s", kafkaBroker, kafkaTopic, kafkaGroupID)

	consumer := kafka.NewKafkaConsumer(kafkaBroker, kafkaTopic, kafkaGroupID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigChan
		log.Println("Señal de apagado recibida. Deteniendo Anomaly-Processor...")
		cancel()
	}()

	logHandler := func(entry domain.LogEntry) error {
		log.Printf("Log entry recibido: %+v", entry)
		return nil
	}

	if err := consumer.Start(ctx, logHandler); err != nil {
		log.Fatalf("Error al iniciar el consumidor de Kafka: %v", err)
	}

	if err := consumer.Close(); err != nil {
		log.Fatalf("Error al cerrar el consumidor de Kafka: %v", err)
	}

	log.Println("Anomaly-Processor finalizado.")
}
