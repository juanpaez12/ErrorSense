# 🔍 ErrorSense: Sistema de Detección de Anomalías en Logs

<div align="center">

![Go Logo](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Kafka Logo](https://img.shields.io/badge/Apache%20Kafka-2.x-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white)
![Docker Logo](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Architecture](https://img.shields.io/badge/Architecture-Microservices-orange?style=for-the-badge)
![Flutter](https://img.shields.io/badge/Flutter-02569B?style=for-the-badge&logo=flutter&logoColor=white)
![Status](https://img.shields.io/badge/Status-En%20Desarrollo-yellow?style=for-the-badge)

**Sistema distribuido de monitoreo en tiempo real para detección proactiva de anomalías en logs**

[🚀 Instalación](#️-configuración-del-entorno-de-desarrollo) • [📖 Documentación](#-arquitectura-del-sistema) • [🤝 Contribuir](#-contribuciones)

</div>

---

## 📋 Tabla de Contenidos

- [🚀 Visión General](#-visión-general-del-proyecto)
- [✨ Características Principales](#-características-principales)
- [🏗️ Arquitectura del Sistema](#️-arquitectura-del-sistema)
- [🛠️ Stack Tecnológico](#️-stack-tecnológico)
- [⚙️ Configuración del Entorno](#️-configuración-del-entorno-de-desarrollo)
- [📦 Microservicios](#-microservicios)
- [🐳 Docker & Orquestación](#-docker--orquestación)
- [🔧 Variables de Entorno](#-configuración-de-variables-de-entorno)
- [🚦 Estado del Proyecto](#-estado-del-proyecto)
- [🤝 Contribuciones](#-contribuciones)

---

## 🚀 Visión General del Proyecto

**ErrorSense** es un sistema de detección de anomalías en logs diseñado como un conjunto de **microservicios** interconectados. Su objetivo principal es monitorear flujos de logs en tiempo real, identificar patrones inusuales o errores críticos, y alertar sobre posibles incidentes operativos.

### 🎯 Problema que Resuelve

En entornos de producción, los sistemas generan miles o millones de logs por segundo. ErrorSense actúa como una herramienta proactiva de vigilancia digital que detecta:

- **Problemas Operacionales**: Picos súbitos de errores, latencia inusual, servicios caídos
- **Amenazas de Seguridad**: Múltiples intentos de login fallidos, accesos desde IPs sospechosas
- **Comportamientos Anómalos**: Patrones que se desvían significativamente de lo "normal"

### 🏛️ Principios de Diseño

- **Arquitectura Hexagonal (Ports & Adapters)**
- **Microservicios Desacoplados**
- **Procesamiento de Eventos Asíncronos**
- **Escalabilidad Horizontal**
- **Observabilidad Completa**

---

## ✨ Características Principales

| Característica | Descripción | Estado |
|---|---|---|
| 🔄 **Procesamiento en Tiempo Real** | Análisis de logs con latencia mínima | ✅ Implementado |
| 🎯 **Detección Inteligente** | Algoritmos de detección de anomalías | 🚧 En desarrollo |
| 📱 **Dashboard Móvil** | Interfaz Flutter para monitoreo | 📋 Planificado |
| 🔔 **Alertas Proactivas** | Notificaciones instantáneas | 📋 Planificado |
| 📊 **Métricas Históricas** | Análisis de tendencias temporales | 📋 Planificado |
| 🐳 **Containerización** | Despliegue con Docker | ✅ Implementado |

---

## 🏗️ Arquitectura del Sistema

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Log Generator │───▶│   Apache Kafka  │───▶│ Anomaly Processor│
│      (Go)       │    │   (Message Bus) │    │      (Go)       │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │                        │
                                ▼                        ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│Flutter Dashboard│    │   ClickHouse    │    │   WebSocket     │
│   (Mobile App)  │◀───│   (Database)    │◀───│   (Real-time)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Componentes del Sistema

1. **🏭 Log-Generator**: Simulador de logs con patrones normales y anómalos
2. **⚡ Apache Kafka**: Bus de mensajes para streaming de eventos
3. **🧠 Anomaly-Processor**: Motor de detección de anomalías
4. **📊 ClickHouse**: Base de datos de series temporales
5. **📱 Flutter Dashboard**: Interfaz móvil de monitoreo
6. **🔗 WebSocket**: Comunicación en tiempo real

---

## 🛠️ Stack Tecnológico

### Backend
- **Go 1.20+**: Servicios de alto rendimiento
- **Apache Kafka**: Streaming de eventos
- **ClickHouse**: Base de datos columnar
- **WebSockets**: Comunicación tiempo real

### Frontend
- **Flutter**: Aplicación móvil multiplataforma
- **Dart**: Lenguaje de programación

### Infraestructura
- **Docker & Docker Compose**: Containerización
- **Apache Zookeeper**: Coordinación de Kafka
- **Terraform**: Infrastructure as Code (futuro)

### Herramientas de Desarrollo
- **GoLand/VS Code**: IDEs recomendados
- **Git**: Control de versiones
- **GitHub Actions**: CI/CD (futuro)

---

## ⚙️ Configuración del Entorno de Desarrollo

### 📋 Requisitos Previos

Asegúrate de tener instalado:

- **Go**: [Versión 1.20 o superior](https://golang.org/doc/install)
- **Docker Desktop**: [Descargar](https://www.docker.com/products/docker-desktop)
- **Git**: [Instalar](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **Flutter** (opcional): Para el dashboard móvil

### 🔧 Instalación

1. **Clonar el Repositorio**
   ```bash
   git clone https://github.com/TU_USUARIO/ErrorSense.git
   cd ErrorSense
   ```

2. **Levantar Infraestructura**
   ```bash
   docker compose up -d
   ```

3. **Verificar Servicios**
   ```bash
   docker compose ps
   ```

> ⏳ **Nota**: Espera 2-3 minutos para que Kafka y Zookeeper se inicien completamente.

---

## 📦 Microservicios

### 1. 🏭 Log-Generator

Simula la generación de logs de una aplicación, enviándolos a Kafka.

#### Configuración

```bash
cd log-generator
```

#### Variables de Entorno (.env)
```env
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=app-logs
```

#### Instalación de Dependencias
```bash
go mod tidy
go get github.com/joho/godotenv
go get github.com/segmentio/kafka-go@v0.4.48
```

#### Ejecución
```bash
go run ./cmd/main.go
```

#### ✅ Verificación
```bash
# Consumir logs desde Kafka
docker exec -it kafka kafka-console-consumer \
  --bootstrap-server localhost:9092 \
  --topic app-logs \
  --from-beginning
```

> 💡 **Tip**: Si descomenta el servicio `kafka-ui` en docker-compose.yml, podrá acceder a una interfaz web en `http://localhost:8080` para monitorear Kafka visualmente.

### 2. 🧠 Anomaly-Processor

> 🚧 **En Desarrollo**: Consumirá logs de Kafka y aplicará lógica de detección.

#### Configuración Futura
```bash
cd anomaly-processor
```

#### Variables de Entorno (.env)
```env
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=app-logs
KAFKA_GROUP_ID=anomaly-processor-group
```

---

## 🐳 Docker & Orquestación

### docker-compose.yml

```yaml
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    container_name: zookeeper
    ports:
      - "2181:2181"
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    networks:
      - app_network

  kafka:
    image: confluentinc/cp-kafka:latest
    container_name: kafka
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
      - "29092:29092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "true"
    networks:
      - app_network

  # Opcional: Kafka UI para monitoreo visual
  # kafka-ui:
  #   image: provectuslabs/kafka-ui:latest
  #   container_name: kafka-ui
  #   depends_on:
  #     - kafka
  #   ports:
  #     - "8080:8080"
  #   environment:
  #     KAFKA_CLUSTERS_0_NAME: local
  #     KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka:9092
  #     KAFKA_CLUSTERS_0_ZOOKEEPER: zookeeper:2181
  #   networks:
  #     - app_network

networks:
  app_network:
    driver: bridge
```

---

## 🔧 Configuración de Variables de Entorno

### .gitignore
```gitignore
# Environment variables
.env

# GoLand / IntelliJ IDEA
.idea/
*.iml
*.ipr
*.iws

# Go binaries and artifacts
*.exe
*.dll
*.so
*.dylib
*.a
*.out
main
coverage.out
build/
dist/

# Dependency directories
vendor/

# IDE specific files
.vscode/
.DS_Store
```

> ⚠️ **Importante**: Los archivos `.env` contienen información sensible y NO deben versionarse.

---

## 🚦 Estado del Proyecto

### ✅ Completado
- [x] Configuración de infraestructura (Docker + Kafka)
- [x] Log Generator con simulación de anomalías
- [x] Integración Kafka para streaming de logs
- [x] Documentación inicial

### 🚧 En Desarrollo
- [ ] Anomaly Processor (detección de patrones)
- [ ] Algoritmos de detección de anomalías
- [ ] Base de datos ClickHouse
- [ ] API WebSocket para tiempo real

### 📋 Planificado
- [ ] Dashboard Flutter
- [ ] Sistema de alertas
- [ ] Métricas y dashboards
- [ ] Despliegue en cloud (Terraform)
- [ ] CI/CD Pipeline
- [ ] Tests automatizados

---

## 🔒 Consideraciones de Seguridad

### 🏠 Configuración de Desarrollo Local

Este proyecto está configurado para **desarrollo local** con los siguientes aspectos de seguridad considerados:

#### ✅ Configuración Actual (Segura para Desarrollo)
- **Puertos expuestos**: Los puertos (9092, 29092, 2181) son configuración de infraestructura, no datos sensibles
- **Sin autenticación**: Kafka corre sin autenticación para simplificar el desarrollo local
- **Red interna**: Los servicios se comunican a través de una red Docker aislada

#### 🏭 Consideraciones para Producción

Para un entorno de producción, se implementarían las siguientes medidas adicionales:

```yaml
# docker-compose.prod.yml (ejemplo)
services:
  kafka:
    ports:
      - "${KAFKA_EXTERNAL_PORT:-9092}:9092"
    environment:
      # Autenticación SASL
      KAFKA_SASL_ENABLED_MECHANISMS: PLAIN
      KAFKA_SASL_MECHANISM_INTER_BROKER_PROTOCOL: PLAIN
      # SSL/TLS
      KAFKA_SSL_ENDPOINT_IDENTIFICATION_ALGORITHM: " "
      KAFKA_SSL_CLIENT_AUTH: requested
    volumes:
      - ./secrets/kafka.keystore.jks:/etc/kafka/secrets/kafka.keystore.jks
      - ./secrets/kafka.truststore.jks:/etc/kafka/secrets/kafka.truststore.jks
```

#### 🔐 Variables de Entorno Sensibles

**Nunca versionar en Git:**
```bash
# .env (excluido por .gitignore)
KAFKA_SASL_JAAS_CONFIG=org.apache.kafka.common.security.plain.PlainLoginModule required username="admin" password="secretpassword";
DATABASE_URL=postgresql://user:password@host:port/db
API_SECRET_KEY=your-secret-key-here
```

#### 🛡️ Mejores Prácticas Implementadas

- ✅ **Separation of Concerns**: Configuración separada por entorno
- ✅ **Secrets Management**: Variables sensibles en archivos .env no versionados  
- ✅ **Network Isolation**: Servicios en red Docker privada
- ✅ **Minimal Exposure**: Solo puertos necesarios expuestos al host

> 💡 **Nota**: Esta configuración prioriza la **facilidad de setup** para evaluación del proyecto. En producción, se implementarían controles de seguridad adicionales según los requisitos específicos del entorno.

---

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Para contribuir:

1. **Fork** del repositorio
2. **Crear** una rama feature (`git checkout -b feature/nueva-funcionalidad`)
3. **Commit** cambios (`git commit -m 'feat: añade nueva funcionalidad'`)
4. **Push** a la rama (`git push origin feature/nueva-funcionalidad`)
5. **Crear** Pull Request

### 📝 Convenciones de Commits

Este proyecto sigue [Conventional Commits](https://www.conventionalcommits.org/) para mantener un historial claro y profesional:

```bash
# Estructura
<tipo>[scope opcional]: <descripción>

# Ejemplos del proyecto
feat(log-generator): Implement anomaly burst injection
fix(kafka): Resolve connection timeout issues  
docs(readme): Add security considerations section
refactor(anomaly-processor): Extract detection algorithms
test(integration): Add Kafka consumer tests
ci: Add GitHub Actions workflow
```
---
<div align="center">

**Construido con ❤️ usando Go, Kafka y Flutter**

[⬆ Volver arriba](#-errorsense-sistema-de-detección-de-anomalías-en-logs)

</div>