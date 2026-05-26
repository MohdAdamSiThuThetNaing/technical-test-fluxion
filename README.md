# Fluxion AI-Driven assessment

## I am using Macos so must be working CPU.

A Docker-based background worker service using MongoDB, RabbitMQ, and Python workers for processing logs, OCR tasks, and background jobs.

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
├── Dockerfile
├── Makefile
├── README.md
├── api
├── cmd
│   ├── api
│   └── worker
├── docker-compose.yml
├── docs
│   └── DOCUMENT.md
├── go.mod
├── go.sum
├── internal
│   ├── ai
│   ├── audit
│   ├── auth
│   ├── dashboard
│   ├── db
│   ├── guard
│   ├── logs
│   ├── migration
│   ├── queue
│   └── users
├── templates
│   ├── ai-suggestion
│   ├── dashboard
│   ├── layouts
│   ├── login.html
│   ├── logs
│   ├── tests
│   └── users
├── test-results.json
└── tests
    ├── ai_test.go
    ├── handler.go
    ├── handler_test.go
    ├── logs_test.go
    ├── queue_test.go
    ├── repository_test.go
    ├── routes.go
    └── user_service_test.go
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

---

# Architecture Overview

```text
Client Browser
      │
      ▼
Gin API Server
      │
      ├── PostgreSQL (Users)
      │
      ├── RabbitMQ Queue
      │         │
      │         ▼
      │    Worker Service
      │         │
      │         ▼
      │     MongoDB Logs
      │
      └── Ollama AI Summary
```

---

# Clean Architecture Layers

```text
handlers
    ↓
services
    ↓
repositories
    ↓
database
```

---

# Implemented Design Patterns

- Repository Pattern
- Service Layer Pattern
- Middleware Authentication
- Worker Queue Architecture
- Dependency Separation
- Route Grouping

---

# Security Features

- Session-based authentication
- Protected admin routes
- Email uniqueness validation
- Password hashing using bcrypt
- Environment variable configuration
- Docker isolated services

---

# AI Integration

The project uses Ollama with the Phi3 model for AI-powered log summarization.

AI features include:

- User activity summarization
- Admin action analysis
- Log event summarization
- Local AI inference without cloud dependency

---

# Background Worker Workflow

```text
User Action
    ↓
RabbitMQ Publish
    ↓
Worker Consume
    ↓
MongoDB Log Storage
    ↓
AI Summary Generation
```

---

# Test Coverage

Current tests include:

- AI integration test
- Queue payload test
- HTTP handler test
- User service/model test

Run tests:

```bash
go test ./... -v
```

Generate coverage:

```bash
go test ./... -cover
```

---

# Docker Containers

| Container Name   | Purpose                 |
| ---------------- | ----------------------- |
| fluxion_api      | Main API server         |
| fluxion_worker   | Background worker       |
| fluxion_postgres | PostgreSQL database     |
| fluxion_mongodb  | MongoDB logs database   |
| fluxion_rabbitmq | RabbitMQ message broker |
| fluxion_ollama   | Ollama AI service       |

---

# Future Improvements

- JWT authentication
- Role-based access control (RBAC)
- Swagger/OpenAPI documentation
- WebSocket real-time logs
- Kubernetes deployment
- CI/CD pipeline
- Redis caching
- Advanced AI analytics
- Email notification service

---

# Development Commands

## Rebuild Containers

```bash
docker compose up --build
```

## Rebuild Without Cache

```bash
docker compose build --no-cache
```

## View API Logs Live

```bash
docker logs -f fluxion_api
```

## View Worker Logs Live

```bash
docker logs -f fluxion_worker
```

## Open PostgreSQL Shell

```bash
docker exec -it fluxion_postgres psql -U admin -d users_db
```

---

# API Endpoints

| Method | Endpoint          | Description      |
| ------ | ----------------- | ---------------- |
| GET    | /login            | Login page       |
| POST   | /login            | Admin login      |
| GET    | /logout           | Logout           |
| GET    | /users            | List users       |
| GET    | /users/create     | Create user page |
| POST   | /users/create     | Create user      |
| POST   | /users/edit/:id   | Update user      |
| POST   | /users/delete/:id | Delete user      |
| GET    | /logs             | View logs        |
| GET    | /ai/test          | AI summary       |
| GET    | /tests            | Test dashboard   |

---

# UI Dashboards

| Dashboard      | URL                           |
| -------------- | ----------------------------- |
| Fluxion Admin  | http://localhost:8080         |
| RabbitMQ UI    | http://localhost:15672        |
| AI Summary     | http://localhost:8080/ai/test |
| Test Dashboard | http://localhost:8080/tests   |

---

This project includes Copilot-inspired AI workflow features powered by Ollama.

### Features

- AI-generated user profile suggestions
- AI-powered CRUD activity insights
- MongoDB log analysis
- Structured JSON responses generated by local LLMs

---

### Example

AI User Suggestion

#### Input

```json
{
  "prompt": "Create a senior backend engineer user"
}
```

#### Output

```json
{
  "suggestion": {
    "name": "Alex Johnson",
    "email": "alex.johnson@techcorp.com",
    "role": "Senior Backend Engineer"
  }
}
```
