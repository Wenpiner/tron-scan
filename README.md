# Tron-Scan

[中文](README_zh.md) | English

Tron-Scan is a blockchain scanning service for the TRON network, developed in Go. It provides comprehensive monitoring and analysis capabilities for TRON network transactions and account activities.

## Features

- Real-time TRON network transaction monitoring
- Account balance and transaction history queries
- Smart contract event listening
- Message queue integration (RabbitMQ)
- RESTful API endpoints
- Distributed tracing support (OpenTelemetry)
- Metrics collection (Prometheus)

## Tech Stack

- Go 1.21+
- go-zero microservice framework
- RabbitMQ message queue
- OpenTelemetry distributed tracing
- Prometheus monitoring
- Docker containerization

## Installation

### Prerequisites

- Go 1.21 or higher
- Docker (optional)

### Local Installation

1. Clone the repository
```bash
git clone https://github.com/wenpiner/tron-scan.git
cd tron-scan
```

2. Install dependencies
```bash
go mod download
```

3. Build the project
```bash
make build
```

### Docker Installation

```bash
docker build -t tron-scan .
docker run -d -p 8888:8888 tron-scan
```

## Configuration

Configuration files are located in the `etc` directory and include:

- Service port settings
- TRON node connection information
- RabbitMQ connection settings
- Monitoring and tracing configurations

## Usage

### Starting the Service

```bash
./tron-scan -f etc/config.yaml
```

### API Documentation

Once the service is running, you can access the API documentation at:
```
http://localhost:8888/swagger
```

## Development

### Project Structure

```
.
├── etc/           # Configuration files
├── internal/      # Internal code
├── tron.go        # Main program entry
├── tron.api       # API definition file
├── Dockerfile     # Docker build file
└── Makefile       # Build scripts
```

### Build Commands

- `make build`: Build the project
- `make run`: Run the project
- `make test`: Run tests
- `make clean`: Clean build files

## License

MIT License 