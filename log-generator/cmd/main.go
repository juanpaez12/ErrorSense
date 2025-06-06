package main

import (
	"ErrorSense/log-generator/generator"
	"ErrorSense/log-generator/kafka"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	if kafkaBroker == "" || kafkaTopic == "" {
		log.Fatalf("Las variables de entorno KAFKA_BROKER o KAFKA_TOPIC no están configuradas.")
	}

	log.Printf("Iniciando Log-Generator con Broker: %s, Topic: %s", kafkaBroker, kafkaTopic)

	producer := kafka.NewKafkaProducer(kafkaBroker, kafkaTopic)
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error al cerrar el productor de Kafka: %v", err)
		}
	}()
	log.Printf("Productor de Kafka inicializado para el broker: %s, tópico: %s", kafkaBroker, kafkaTopic)

	logGenerator := generator.NewLogGenerator(producer, kafkaTopic)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go logGenerator.Start()

	<-ctx.Done()
	log.Println("Señal de apagado recibida. Deteniendo Log-Generator...")

	time.Sleep(2 * time.Second)
	log.Println("Cerrando el microservicio log-generator...")
}
