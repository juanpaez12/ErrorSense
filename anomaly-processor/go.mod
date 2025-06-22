module ErrorSense/anomaly-processor

go 1.24.0

require (
	github.com/joho/godotenv v1.5.1
	github.com/segmentio/kafka-go v0.4.48
)

require (
	ErrorSense/log-generator v0.0.0
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
)

replace ErrorSense/log-generator => ../log-generator
