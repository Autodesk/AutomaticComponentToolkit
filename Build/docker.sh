#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

echo "Build Docker image"
docker build -t act-build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) .
