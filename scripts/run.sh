#!/usr/bin/env bash
set -e

echo "Starting backend..."
go run ./cmd/server &
SERVER_PID=$!

cleanup() {
  echo "Stopping backend..."
  kill "$SERVER_PID" 2>/dev/null || true
}
trap cleanup EXIT

echo "Starting proxy..."
go run ./cmd/proxy
