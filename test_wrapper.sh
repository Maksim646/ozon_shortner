#!/bin/bash
set -x

PROJECT_ROOT=$(git rev-parse --show-toplevel)

echo "PROJECT_ROOT: $PROJECT_ROOT"

cd "$PROJECT_ROOT"

TEST_DIR="./internal/api/server/restapi/handler"

echo "TEST_DIR: $TEST_DIR"

go test -v "$TEST_DIR" 2>&1