#!/bin/bash

set -e

host="postgres"
port="5432"
user="postgres"

echo "⏳ Waiting for PostgreSQL at $host:$port to be ready..."

until pg_isready -h "$host" -p "$port" -U "$user" > /dev/null 2>&1; do
  echo "❌ PostgreSQL not ready yet... retrying in 1s"
  sleep 1
done

echo "✅ PostgreSQL is ready. Starting the app..."
exec "$@"
