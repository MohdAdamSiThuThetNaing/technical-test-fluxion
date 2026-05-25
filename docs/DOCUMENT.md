# Senior Software Engineer Assessment Documentation

# Project Name

## Managing User Records App

---

# Overview

This project is an AI-driven user management system developed with Golang using the Gin framework.

The system provides:

- Admin authentication
- User CRUD management
- Asynchronous activity logging
- MongoDB log storage
- RabbitMQ background queue processing
- AI-powered workflow integrations using Ollama
- Dockerized microservice architecture
- Unit testing coverage

The application follows clean architecture principles and implements scalable backend patterns suitable for enterprise-level systems.

---

# Technology Stack

| Technology    | Purpose               |
| ------------- | --------------------- |
| Golang (Gin)  | Backend API Framework |
| PostgreSQL    | Relational Database   |
| MongoDB       | NoSQL Log Storage     |
| RabbitMQ      | Asynchronous Queue    |
| Ollama + Phi3 | AI-Powered Services   |
| Docker        | Containerization      |
| Gin Sessions  | Authentication        |
| bcrypt        | Password Hashing      |

---

# Database Schemas

## Users Schema

| Field     | Type           |
| --------- | -------------- |
| ID        | UUID / Integer |
| Name      | String         |
| Email     | String         |
| Password  | Hashed String  |
| CreatedAt | Timestamp      |
| UpdatedAt | Timestamp      |

---

## UserLogs Schema

| Field     | Type           |
| --------- | -------------- |
| UserID    | UUID / Integer |
| Event     | String         |
| Data      | JSON / String  |
| CreatedAt | Timestamp      |

---

# Requirements Implementation

## 1. Admin Login Functionality

Implemented session-based authentication using Gin sessions.

Features:

- Login page
- Session management
- Protected admin routes
- Logout functionality
- bcrypt password hashing

Endpoints:

```text
GET  /login
POST /login
GET  /logout
```

---

## 2. User Management CRUD

Implemented full CRUD operations for user management.

Features:

- Create users
- List users
- Update users
- Delete users
- Email uniqueness validation

Endpoints:

```text
GET  /users
GET  /users/create
POST /users/create
POST /users/edit/:id
POST /users/delete/:id
```

---

## 3. User List Data Table

Implemented user listing interface with admin dashboard UI.

Features:

- User listing
- Structured table layout
- CRUD action buttons
- Responsive admin templates

---

## 4. Asynchronous Logging Mechanism

Implemented event-driven asynchronous logging architecture.

Workflow:

```text
User CRUD Action
        ↓
RabbitMQ Publish
        ↓
Background Worker Consume
        ↓
MongoDB Log Storage
```

Logged events include:

- USER_CREATED
- USER_UPDATED
- USER_DELETED
- USER_LOGIN

---

## 5. Admin UI

Implemented admin panel using HTML templates with dashboard layouts.

Pages:

- Login page
- User management dashboard
- Logs dashboard
- AI insights dashboard
- Test dashboard

---

## 6. Relational Database Usage

Implemented PostgreSQL for relational user data storage.

Used for:

- User records
- Authentication data
- Session-related data

---

## 7. NoSQL Database Usage

Implemented MongoDB for activity log storage.

Used for:

- CRUD activity logs
- Admin events
- AI log analysis

---

## 8. Unit Testing

Implemented unit testing coverage across affected areas.

Covered modules:

- AI services
- Queue services
- HTTP handlers
- User services
- Repository layer

Run tests:

```bash
go test ./... -v
```

Generate coverage:

```bash
go test ./... -cover
```

---

# AI-Driven Workflows

The project includes AI-powered workflow integrations using Ollama with the Phi3 model.

AI capabilities include:

- CRUD activity summarization
- Admin action analysis
- MongoDB log analysis
- AI-generated user profile suggestions
- Structured JSON AI responses

---

# Copilot-Inspired AI User Suggestions

Implemented AI-assisted user generation workflow.

Endpoint:

```text
POST /ai/user-suggestion
```

Example Request:

```json
{
  "prompt": "Create a senior backend engineer user"
}
```

Example Response:

```json
{
  "suggestion": {
    "name": "Alex Johnson",
    "email": "alex.johnson@techcorp.com",
    "role": "Senior Backend Engineer"
  }
}
```

Workflow:

```text
Admin Prompt
      ↓
Gin API
      ↓
Ollama AI Service
      ↓
Structured JSON Response
      ↓
User Creation Workflow
```

---

# AI CRUD Insights

Implemented AI-powered log analysis endpoint.

Endpoint:

```text
GET /ai/test
```

Example Response:

```json
{
  "result": {
    "total_user_created": 5,
    "total_user_updated": 2,
    "total_user_deleted": 1,
    "latest_admin_action": "USER_UPDATED"
  }
}
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
      └── Ollama AI Services
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

# Design Patterns Implemented

- Repository Pattern
- Service Layer Pattern
- Middleware Authentication
- Worker Queue Architecture
- Dependency Separation
- Route Grouping

---

# Security Features

- Session-based authentication
- Password hashing using bcrypt
- Protected admin routes
- Email uniqueness validation
- Environment variable configuration
- Docker isolated services

---

# Deliverables Completed

## Completed Users Data

Implemented PostgreSQL-backed user management system with CRUD operations.

---

## Completed User Logs Data

Implemented asynchronous MongoDB logging system with RabbitMQ worker processing.

---

## Unit Testing

Implemented testing coverage for core application modules.

---

## AI-Driven Workflows

Implemented Ollama-powered AI workflows including:

- AI CRUD insights
- AI-generated user suggestions
- AI log analysis
- Structured AI JSON responses

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

# API Endpoints

| Method | Endpoint            | Description        |
| ------ | ------------------- | ------------------ |
| GET    | /login              | Login page         |
| POST   | /login              | Admin login        |
| GET    | /logout             | Logout             |
| GET    | /users              | List users         |
| GET    | /users/create       | Create user page   |
| POST   | /users/create       | Create user        |
| POST   | /users/edit/:id     | Update user        |
| POST   | /users/delete/:id   | Delete user        |
| GET    | /logs               | View logs          |
| GET    | /ai/test            | AI summary         |
| POST   | /ai/user-suggestion | AI user suggestion |
| GET    | /tests              | Test dashboard     |

---

# Conclusion

This assessment demonstrates the implementation of a scalable AI-driven CRUD management system using modern backend engineering practices.

The project combines:

- Clean architecture
- Asynchronous processing
- Relational + NoSQL databases
- AI-powered workflows
- Dockerized infrastructure
- Unit testing
- Secure authentication

to deliver a production-style backend system suitable for enterprise-level development.
