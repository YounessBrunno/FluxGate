#!/usr/bin/env bash
set -e

echo "Starting backend..."
air --build.cmd "go build -o ./tmp/server ./cmd/server" --build.bin "./tmp/server" &
SERVER_PID=$!

cleanup() {
  echo "Stopping backend..."
  kill "$SERVER_PID" 2>/dev/null || true
}
trap cleanup EXIT

echo "Starting proxy..."
air --build.cmd "go build -o ./tmp/proxy ./cmd/proxy" --build.bin "./tmp/proxy"