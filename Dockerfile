########################################
# -------- 1. BUILD STAGE ------------ #
########################################
FROM golang:1.24.4 AS builder

WORKDIR /app

# ✅ Disable CGO to avoid toolchain segfaults in container
ENV CGO_ENABLED=0

# copy and download deps first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the source and build
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o todo-app

########################################
# -------- 2. RUNTIME STAGE ---------- #
########################################
FROM debian:bookworm-slim

WORKDIR /app

# optional: install psql client (for migrations/debugging)
RUN apt-get update \
 && apt-get install -y --no-install-recommends postgresql-client \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

# copy compiled Go binary + wait script
COPY --from=builder /app/todo-app .
COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

# default command: wait for PG, then start the app
CMD ["./wait-for-postgres.sh", "postgres", "5432", "./todo-app"]
