#!/bin/bash
# Use strict bash
set -euo pipefail
IFS=$'\n\t'

# Get the current gopath
GOPATH="$(go env GOPATH)"

cd ${GOPATH}/src/github.com/housekpr/core/srv
CC=arm-linux-gnueabi-gcc CXX=arm-linux-gnueabi-g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -o ../bin/housekpr
echo "[OK] Build done."