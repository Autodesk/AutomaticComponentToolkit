#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Go example"

[ -d build ] && rm -rf build
mkdir build
GO111MODULE=off go build -o build/RTTI_example RTTI_example.go 

echo "Test C++ library"
./build/RTTI_example $PWD/../../Implementations/Cpp/build/rtti$OSLIBEXT

echo "Test Pascal library"
./build/RTTI_example $PWD/../../Implementations/Pascal/build/rtti$OSLIBEXT
