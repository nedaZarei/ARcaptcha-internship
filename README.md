# ARcaptcha Internship Repository

A comprehensive **Go-based project repository** containing algorithms, design patterns, learning notes, and a complete web application developed during my **ARcaptcha internship**.

---

### 📚 Learning Notes
Documentation on software engineering concepts:

- **Backend Communication Patterns**
  - Publish/Subscribe (`publish-subscribe.md`)
  - Push Pattern (`push-pattern.md`)
  - Request-Response (`request-response-pattern.md`)
- **Design Patterns**
  - Adapter (`adapter-pattern-go.md`)
  - Builder (`builder_pattern_go.md`)
  - Factory Method (`factory-method-go.md`)
  - SOLID (`solid-golang.md`, `solid_principles.md`)
- **Testing Strategies**
  - Faking (`faking-in-go.md`)
  - Mocking (`mocking-in-go.md`)
  - Monkey Patching (`monkey-patching-go.md`)

### 🚀 Main Project: Apartment Management System
A full **Go web application** implementing Clean Architecture principles.

**Core Components:**
- `cmd/` → Application entry points (`root.go`, `start.go`) 
- `config/` → Configuration management (`config.go`, `config.example.yml`) 
- `internal/` → Private app code, structured by layers
  - **Models** (`models/`) → Entities: User, Apartment, Bill, Payment, InvitationLink 
  - **DTOs** (`dto/`) → Transfer objects for API communication 
  - **Repositories** (`repositories/`) → Data access layer + mocks + tests 
  - **Services** (`services/`) → Business logic + service tests 
  - **HTTP** (`http/`) → Handlers, middleware, routes, utilities 
  - **App** (`app/`) → DB, Redis, MinIO integrations 
  - **External** → Image processing, Notifications, Payment service 

---

## 🛠️ Technologies Used
- **Language**: Go (Golang) 
- **Architecture**: Clean Architecture + Domain-Driven Design 
- **Database**: PostgreSQL 
- **Cache**: Redis 
- **Storage**: MinIO 
- **API**: RESTful, documented with Swagger/OpenAPI 
- **Testing**: Native Go testing, mocks, coverage 
- **Infra**: Docker + Docker Compose 

---

## 🏗️ Patterns & Practices
- Repository Pattern 
- Service Layer Pattern 
- DTO Pattern 
- Middleware Pattern 
- Mocking & Test Isolation 
- Factory Method 
- Builder Pattern 
- Adapter Pattern 
- SOLID Principles 

---

## 📋 Key Features
1. **User Management** → CRUD + Authentication 
2. **Apartment System** → Multi-tenant apartment management 
3. **Billing** → Automated bill generation & tracking 
4. **Payments** → Secure payment processing 
5. **Invitations** → Shareable invite links 
6. **Image Management** → File upload & processing 
7. **Notifications** → Event-driven notification service 
8. **Testing** → Extensive unit & integration coverage 

---

## 📖 Learning Outcomes
During this internship, I gained practical experience in:

- Advanced **Go programming patterns** 
- **Clean Architecture** & Domain-Driven Design 
- Building and testing **REST APIs** 
- Designing **database schemas** with PostgreSQL 
- Integrating **Redis & MinIO** into backend services
- Applying **design patterns** (Factory, Builder, Adapter, Repository, etc.) 
- Writing **comprehensive tests** with mocks 
- Using **Docker & Docker Compose** for dev environments 

---

## 📜 Documentation
- API: `swagger.yaml` (viewable in [Swagger Editor](https://editor.swagger.io/)) 
- Design Patterns: `notes/design-patterns/` 
- Testing Guide: `notes/testing/` 

---

This repo represents my learning journey during the ARcaptcha internship. 
