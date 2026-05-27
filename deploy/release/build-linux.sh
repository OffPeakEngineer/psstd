#!/usr/bin/env sh
set -eu

mkdir -p dist
version="${PULSED_VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo dev)}"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -X main.appVersion=${version}" -o dist/pulsed-linux-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w -X main.appVersion=${version}" -o dist/pulsed-linux-arm64 .
