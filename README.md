# 📝 TODO API – Go + PostgreSQL + Redis + Docker

A production-ready TODO API built using **Go**, backed by **PostgreSQL** for persistent storage, and **Redis** for in-memory caching.

This project is fully containerized using **Docker** and managed via **Docker Compose**, making it easy to run locally or deploy to cloud platforms like **AWS EC2**.

---

## 📦 Features

- RESTful API with CRUD operations for TODO items
- PostgreSQL integration using `pgxpool`
- Redis caching for optimized read performance (`GET /todos`)
- Graceful handling of database connections and caching
- Clean code structure: separated into handler, service, db, and cache layers
- Dockerized with multi-stage build
- Environment configurable
- Easy to test with Postman or curl

---

## 🏗️ Tech Stack

| Layer        | Technology          |
|--------------|----------------------|
| Language     | Go 1.24.x            |
| Database     | PostgreSQL 16        |
| Cache        | Redis 7              |
| ORM/Driver   | `pgx/v5`             |
| Docker       | Multi-stage Dockerfile |
| Dev Tools    | Docker Compose       |

---

## 🚀 Quick Start with Docker

> Make sure you have **Docker** and **Docker Compose** installed.

### 🐳 1. Clone the repository

```bash
git clone https://github.com/sanyam-harness/todo-api-go-postgre-redis-docker.git
cd todo-api-go-postgre-redis-docker
````

### 🐳 2. Build and start the services

```bash
docker compose up --build
```

You should see logs like:

```
✅ PostgreSQL is ready. Starting the app...
✅ Connected to PostgreSQL using pgxpool
✅ Connected to Redis successfully: PONG
🚀 Server running at http://localhost:8080
```

---

## 🧪 API Testing Guide

Once the app is running at **[http://localhost:8080](http://localhost:8080)**, use **Postman** or `curl` to test the following endpoints.

### 📌 Base URL

```
http://localhost:8080
```

---

### ✅ Create a TODO

* **POST** `/todos`
* **Body (JSON)**:

```json
{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread"
}
```

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries", "description":"Milk, eggs, bread"}'
```

---

### 📄 List all TODOs

* **GET** `/todos`

```bash
curl http://localhost:8080/todos
```

> ℹ️ This endpoint uses **Redis caching** to boost performance.

---

### 🔍 Get a TODO by ID

* **GET** `/todos/{id}`

```bash
curl http://localhost:8080/todos/1
```

---

### ✏️ Update a TODO

* **PUT** `/todos/{id}`
* **Body**:

```json
{
  "title": "Buy groceries and fruits",
  "description": "Milk, eggs, apples"
}
```

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries and fruits","description":"Milk, eggs, apples"}'
```

---

### 🗑️ Delete a TODO

* **DELETE** `/todos/{id}`

```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## 🧠 Redis Caching (How it works)

* When `GET /todos` is called:

  * The app first checks Redis for a cached list.
  * If not found, it queries PostgreSQL, then **stores the result in Redis** for next time.

* On any **write/update/delete**, the cache is **invalidated**.

---

## 🔧 Project Structure

```
todo-api-go-postgre-redis-docker/
├── main.go              # App entry point
├── db.go                # PostgreSQL connection
├── cache.go             # Redis connection + logic
├── handler.go           # HTTP handlers
├── service.go           # Business logic layer
├── models.go            # Data models
├── wait-for-postgres.sh # Script to delay app until PG is ready
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Orchestrates app + db + redis
└── README.md            # Project docs
```

---

## 🧪 Running Without Docker (optional)

If you'd rather run the Go app natively:

```bash
export DATABASE_URL=postgres://postgres:admin123@localhost:5432/tododb
export REDIS_ADDR=localhost:6379
go run main.go
```

Make sure PostgreSQL and Redis are running locally.

---

## 🧑‍💻 Contributing (optional)

Pull requests are welcome! To contribute:

```bash
git checkout -b feature/my-feature
# make changes
git commit -m "Add my feature"
git push origin feature/my-feature
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 🙌 Author

Developed by [Sanyam Jain](https://github.com/sanyam-harness)

