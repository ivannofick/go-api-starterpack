# go-api-starterpack

A lightweight and opinionated Go API starter pack featuring a clean project structure, PostgreSQL integration, JWT authentication, and transaction-safe CRUD operations.

This project is designed to be **simple enough for learning**, yet **structured enough for real-world usage**.

---

## ✨ Features

* Go (`net/http`) REST API
* PostgreSQL with GORM
* JWT Authentication
* Transaction-safe operations
* Centralized API response helper
* Try–Catch style error handling (`panic` + `recover`)
* Clean and scalable folder structure
* Docker support for local PostgreSQL

---

## 📁 Project Structure (Honest Review)

The current structure is intentionally kept **lean and practical** — no over-engineering.

```
app/
├── controllers      # ✅ Clear HTTP layer (request / response)
├── middleware       # ✅ Auth middleware (recover can be added later)
├── models           # ✅ Database models (GORM)
├── services         # ✅ Reusable logic (JWT for now)

config/              # ✅ Infrastructure configuration

database/            # ✅ Database migration

dto/                 # ✅ Request / Response data shapes

libs/                # ✅ Infrastructure helpers
│   ├── response.go  # Standard API response
│   ├── trycatch.go  # Panic-safe error handling
│   └── with_transaction.go # Transaction wrapper

routes/               # ✅ API routing

deployments/          # ✅ Docker & infrastructure setup
```

### Why this structure?

* Controllers stay **thin** and focused
* Business logic is reusable
* Infrastructure concerns are isolated
* Easy to extend when complexity grows

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-api-starterpack.git
cd go-api-starterpack
```

### 2. Setup environment variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=go_crud
JWT_SECRET=supersecretkey
```

### 3. Run PostgreSQL (Docker)

```bash
docker-compose -f deployments/docker-compose-postgres.yml up -d
```

### 4. Run the application

```bash
go mod tidy
go run main.go
```

Server will run at:

```
http://localhost:8000
```

---

## 🔐 Authentication Flow

1. Register user
2. Login to receive JWT token
3. Use token in `Authorization: Bearer <token>` header
4. Access protected routes

---

## 📌 Design Principles

* **Structure follows complexity** — not the other way around
* Avoid premature abstraction
* Keep controllers clean
* Centralize cross-cutting concerns (auth, response, transaction)

---

## 🧭 When to Extend This Project

Add more layers (repository, interfaces, services) when:

* Business logic grows
* Multiple models interact
* Unit testing becomes critical
* You need to swap database or ORM

---

## 🛠️ Tech Stack

* Go
* PostgreSQL
* GORM
* JWT
* Docker

---

## 📄 License

MIT License

---

Happy hacking! 🚀
