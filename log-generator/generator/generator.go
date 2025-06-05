package generator

import (
	"ErrorSense/log-generator/domain"
	"ErrorSense/log-generator/kafka"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type LogGenerator struct {
	producer kafka.LogProducer
	topic    string
	stopChan chan struct{}
}

func NewLogGenerator(producer kafka.LogProducer, topic string) *LogGenerator {
	return &LogGenerator{
		producer: producer,
		topic:    topic,
		stopChan: make(chan struct{}),
	}
}

func (lg *LogGenerator) Start() {
	log.Println("Iniciando generador de logs...")
	ticker := time.NewTicker(200 * time.Microsecond)
	defer ticker.Stop()

	anomalyInjectionInterval := 2 * time.Minute
	lastAnomalyTime := time.Now()

	for {
		select {
		case <-lg.stopChan:
			log.Println("Deteniendo generador de logs...")
			return
		case <-ticker.C:
			if time.Since(lastAnomalyTime) > anomalyInjectionInterval {
				lg.generateAnomalyLogBurst()
				lastAnomalyTime = time.Now()
			} else {
				lg.generateNormalLog()
			}
		}
	}
}

func (lg *LogGenerator) Stop() {
	close(lg.stopChan)
}

func (lg *LogGenerator) generateNormalLog() {
	logEntry := domain.LogEntry{
		Timestamp: time.Now(),
		Source:    getRandomSource(),
		UserID:    fmt.Sprintf("user_%d", rand.Intn(1000)), // 1000 usuarios simulados
		IPAddress: generateRandomIP(),
		SessionID: generateRandomString(10),
	}

	roll := rand.Float64()
	if roll < 0.8 {
		logEntry.Level = "INFO"
		logEntry.Message = getRandomInfoMessage(logEntry.Source)
		logEntry.Endpoint = getRandomEndpoint(logEntry.Source)
		logEntry.StatusCode = 200
	} else if roll < 0.95 {
		logEntry.Level = "WARN"
		logEntry.Message = getRandomWarnMessage(logEntry.Source)
		logEntry.Endpoint = getRandomEndpoint(logEntry.Source)
		logEntry.StatusCode = 400 + rand.Intn(3) // 400, 401, 402, etc.
	} else {
		logEntry.Level = "ERROR"
		logEntry.Message = getRandomErrorMessage(logEntry.Source)
		logEntry.Endpoint = getRandomEndpoint(logEntry.Source)
		logEntry.StatusCode = 500 + rand.Intn(3) // 500, 501, 502, etc.
	}

	err := lg.producer.Produce(logEntry)
	if err != nil {
		log.Printf("Error al escribir log en Kafka: %v\n", err)
	}
}

func (lg *LogGenerator) generateAnomalyLogBurst() {
	log.Println("------ INYECTANDO RÁFAGA DE ANOMALÍA (FUERZA BRUTA) ------")
	anomalyUserID := fmt.Sprintf("user_admin_%d", rand.Intn(5)) // Simula intentos contra pocas cuentas admin
	anomalyIP := generateRandomIP()                             // La IP puede variar o ser la misma, aquí la hacemos fija para la ráfaga
	anomalySessionID := generateRandomString(15)

	burstSize := 50 + rand.Intn(50)

	for i := 0; i < burstSize; i++ {
		logEntry := domain.LogEntry{
			Timestamp:  time.Now(),
			Level:      "WARN",
			Message:    "Login attempt failed: invalid credentials",
			Source:     "auth-service",
			UserID:     anomalyUserID,
			IPAddress:  anomalyIP,
			Endpoint:   "/auth/login",
			StatusCode: 401,
			SessionID:  anomalySessionID,
		}
		err := lg.producer.Produce(logEntry)
		if err != nil {
			log.Printf("Error al producir log anómalo %d: %v", i, err)
		}
		time.Sleep(10 * time.Millisecond) // Pequeña pausa para simular ráfaga
	}
	log.Println("------ RÁFAGA DE ANOMALÍA TERMINADA ------")
}

// --- Funciones auxiliares para generar datos aleatorios ---

var sources = []string{"web-app", "auth-service", "payment-api", "user-profile-service", "notification-service"}
var infoMessages = map[string][]string{
	"web-app":              {"Page loaded", "Item added to cart", "User navigated"},
	"auth-service":         {"User logged in successfully", "Token validated"},
	"payment-api":          {"Transaction initiated", "Payment request sent"},
	"user-profile-service": {"Profile viewed", "Settings updated"},
	"notification-service": {"Email sent", "SMS delivered"},
}
var warnMessages = map[string][]string{
	"web-app":              {"Slow response time", "Missing required field"},
	"auth-service":         {"Invalid session token", "Account locked temporarily"},
	"payment-api":          {"Payment gateway timeout", "Insufficient funds"},
	"user-profile-service": {"Invalid profile data", "Data not found"},
	"notification-service": {"SMS delivery failed", "Email sending failed"},
}
var errorMessages = map[string][]string{
	"web-app":              {"Internal server error", "Database connection lost"},
	"auth-service":         {"Critical authentication error", "Database unavailable"},
	"payment-api":          {"Failed to process payment", "Gateway connection failed"},
	"user-profile-service": {"Failed to load profile", "User deletion failed"},
	"notification-service": {"Unhandled exception", "External service unreachable"},
}
var endpoints = map[string][]string{
	"web-app":              {"/", "/products", "/cart", "/checkout"},
	"auth-service":         {"/auth/login", "/auth/register", "/auth/token"},
	"payment-api":          {"/payment/process", "/payment/status"},
	"user-profile-service": {"/users/{id}", "/users/settings"},
	"notification-service": {"/notify/email", "/notify/sms"},
}

func getRandomSource() string {
	return sources[rand.Intn(len(sources))]
}

func getRandomInfoMessage(source string) string {
	msgs := infoMessages[source]
	if len(msgs) == 0 {
		return "Generic INFO message"
	}
	return msgs[rand.Intn(len(msgs))]
}

func getRandomWarnMessage(source string) string {
	msgs := warnMessages[source]
	if len(msgs) == 0 {
		return "Generic WARN message"
	}
	return msgs[rand.Intn(len(msgs))]
}

func getRandomErrorMessage(source string) string {
	msgs := errorMessages[source]
	if len(msgs) == 0 {
		return "Generic ERROR message"
	}
	return msgs[rand.Intn(len(msgs))]
}

func getRandomEndpoint(source string) string {
	msgs := endpoints[source]
	if len(msgs) == 0 {
		return "/generic/path"
	}
	return msgs[rand.Intn(len(msgs))]
}

func generateRandomIP() string {
	return fmt.Sprintf("192.168.%d.%d", rand.Intn(256), rand.Intn(256))
}

func generateRandomString(Length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, Length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
