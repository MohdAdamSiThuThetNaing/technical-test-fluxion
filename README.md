# Fluxion Worker Service

A Docker-based background worker service using MongoDB, RabbitMQ, and Python workers for processing logs, OCR tasks, and background jobs.

---

# Prerequisites

Before starting the project, make sure the following are installed:

- Docker
- Docker Compose
- Git
- Go (for running tests)

---

# Clone the Repository

```bash
git clone <your-repository-url>
cd <project-folder>
```

---

# Start the Services

Build and run all containers:

```bash
docker compose up --build -d
```

Check running containers:

```bash
docker ps -a
```

---

# Docker Commands

## View Running Containers

```bash
docker ps -a
```

## View Worker Logs

```bash
docker logs fluxion_worker
```

## Open MongoDB Shell

```bash
docker exec -it fluxion_mongodb mongosh
```

Inside MongoDB shell:

```javascript
use fluxion_logs
show collections
db.logs.find().pretty()
```

## Restart Containers

```bash
docker compose restart
```

## Stop Containers

```bash
docker compose down
```

## Stop Containers and Remove Volumes

```bash
docker compose down -v
```

## Remove Unused Docker Volumes

```bash
docker volume prune -f
```

---

# RabbitMQ Dashboard

RabbitMQ Management UI:

```text
http://localhost:15672/#/
```

Default credentials:

```text
Username: guest
Password: guest
```

---

# Run Tests

Execute all Go tests with verbose output:

```bash
go test ./... -v
```

---

# Project Structure

```text
.
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── docs/
├── internal/
│   ├── ai/
│   ├── auth/
│   ├── db/
│   ├── guard/
│   ├── logs/
│   ├── queue/
│   └── users/
├── templates/
├── tests/
│   ├── ai_test.go
│   ├── queue_test.go
│   └── user_service_test.go
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md

```

---

# Troubleshooting

## Worker Container Not Running

Check logs:

```bash
docker logs fluxion_worker
```

## MongoDB Connection Issues

Verify MongoDB container:

```bash
docker ps -a
```

Enter Mongo shell:

```bash
docker exec -it fluxion_mongodb mongosh
```

## RabbitMQ Access Issues

Ensure RabbitMQ container is running and port `15672` is exposed.

---

# Development Notes

- MongoDB database name: `fluxion_logs`
- Worker container name: `fluxion_worker`
- MongoDB container name: `fluxion_mongodb`
- RabbitMQ default management port: `15672`

---

# Useful Commands Summary

```bash
# Start services
docker compose up --build -d

# Check containers
docker ps -a

# Worker logs
docker logs fluxion_worker

# Mongo shell
docker exec -it fluxion_mongodb mongosh

# Run tests
go test ./... -v

# Stop services
docker compose down
```
