package main

import (
	"ErrorSense/log-generator/generator"
	"ErrorSense/log-generator/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	kafkaBroker = "localhost:9092"
	KafkaTopic  = "app-logs"
)

func main() {
	log.Println("Iniciando el microservicio log-generator...")
	producer := kafka.NewKafkaProducer(kafkaBroker, KafkaTopic)
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error al cerrar el productor de Kafka: %v", err)
		}
	}()
	log.Printf("Productor de Kafka inicializado para el broker: %s, tópico: %s", kafkaBroker, KafkaTopic)

	logGenerator := generator.NewLogGenerator(producer, KafkaTopic)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go logGenerator.Start()

	sig := <-stopChan
	log.Printf("Recibiendo señal %s, cerrando el generador de logs...", sig)
	logGenerator.Stop()
	time.Sleep(2 * time.Second)
	log.Println("Cerrando el microservicio log-generator...")
}
