#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Go example"

[ -d build ] && rm -rf build
mkdir build
GO111MODULE=off go build -o build/LibPrimes_example LibPrimes_example.go 

echo "Test C++ library"
./build/LibPrimes_example $PWD/../../Implementations/Cpp/build/libprimes$OSLIBEXT

echo "Test Pascal library"
./build/LibPrimes_example $PWD/../../Implementations/Pascal/build/libprimes$OSLIBEXT
