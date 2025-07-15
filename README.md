# 🚀 GoLang Learning Series

Welcome to the **Golang Learning Series** – a complete journey from Go basics to building powerful, real-world backend applications using Go, MongoDB, concurrency patterns, and more.

---

## 📘 Overview

This repository is a structured collection of everything I’ve learned while diving into the Go programming language. It includes concepts, code samples, and real-world use cases like building REST APIs, working with MongoDB, handling concurrency, and avoiding race conditions.

---

## 🗂️ Topics Covered

### ✅ 1. Go Modules & Project Setup
- Initialized modules with `go mod init`
- Managed dependencies with `go mod tidy`
- Understood `go.mod` and `go.sum` files

### 🌐 2. CRUD REST API with net/http
- Built a full REST API with native `net/http` and `gorilla/mux`
- Routes included:
  - `GET /courses`
  - `GET /course/{id}`
  - `POST /course`
  - `PUT /course/{id}`
  - `DELETE /course/{id}`
- Learned JSON encoding/decoding and request handling

### 🍃 3. MongoDB Integration
- Connected to MongoDB using the official Go driver
- Stored data in a `watchlist` collection of a `netflix` database
- Created a centralized MongoDB config with `init()` connection
- Shared `collection` instance across controllers using exported variables

### 🧱 4. MVC Project Structure
- Structured code into:
  - `models/` (e.g., `userModel.go`)
  - `controllers/` (e.g., `userController.go`)
  - `config/` (for DB connections)
- Used imports like `import "yourmodule/models"` to share code

### 📦 5. BSON in MongoDB
- `bson.M` (unordered map) for quick operations
- `bson.D` (ordered list) for queries where order matters

### 🧠 6. Concurrency in Go
- Learned the concept of **concurrency** (multiple tasks interleaved)
- Used `go` keyword to create **goroutines**
- Controlled execution with `time.Sleep()` and `WaitGroup`

### 🔄 7. Goroutines
- Lightweight threads managed by the Go runtime
- Syntax: `go someFunction()`
- Asynchronous function calls without blocking the main thread

### ⏳ 8. WaitGroup (sync package)
- `sync.WaitGroup` used to wait for all goroutines to finish
- Methods:
  - `Add(n)`
  - `Done()`
  - `Wait()`

### 🔒 9. Mutex
- Used `sync.Mutex` to avoid data races
- Locking and unlocking critical sections to ensure safe access

### ⚔️ 10. Race Conditions
- Explored how unsynchronized goroutines cause race conditions
- Demonstrated with shared counters and improper synchronization
- Solved with Mutex

### 📬 11. Channels
- Channels as pipelines for goroutine communication
- Unbuffered and buffered channels:
  - Send: `ch <- value`
  - Receive: `val := <- ch`
- Demonstrated real-world channel usage

---

## 📂 Example Directory Structure
golang-series/
├── config/
│ └── db.go
├── controllers/
│ └── userController.go
├── models/
│ └── userModel.go
├── main.go
├── go.mod
├── go.sum
└── README.md


---

## 💻 Author

**Sachidanand Sharma**  
Learning Golang, Web Backends, and Distributed Systems  
🔗 [GitHub](https://github.com/SachidanandSharma2162)

---

## 📌 Upcoming Topics

- Select statements
- Context and cancellation
- Middleware integration
- Unit & integration testing in Go
- Dockerizing Go apps
- gRPC microservices

---

> “Don’t fear the Go routine. Fear the unguarded shared memory.” 🧠

Happy Hacking & Keep Learning! 🙌

