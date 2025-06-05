package domain

import "time"

type LogEntry struct {
	Timestamp     time.Time `json:"timestamp"`                // Marca de tiempo del evento
	Level         string    `json:"level"`                    // Nivel del log (INFO, WARN, ERROR, DEBUG, FATAL)
	Message       string    `json:"message"`                  // Descripción del evento
	Source        string    `json:"source"`                   // Origen del log (e.g., "web-app", "auth-service", "payment-api")
	UserID        string    `json:"user_id,omitempty"`        // ID de usuario, opcional
	IPAddress     string    `json:"ip_address,omitempty"`     // Dirección IP del cliente, opcional
	Endpoint      string    `json:"endpoint,omitempty"`       // Endpoint de la petición HTTP/RPC, opcional
	StatusCode    int       `json:"status_code,omitempty"`    // Código de estado HTTP o interno, opcional (e.g., 200, 401, 500)
	SessionID     string    `json:"session_id,omitempty"`     // ID de sesión, opcional
	TransactionID string    `json:"transaction_id,omitempty"` // ID de transacción, opcional
}
