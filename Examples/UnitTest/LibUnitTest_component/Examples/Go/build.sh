#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Go example"

[ -d build ] && rm -rf build
mkdir build
GO111MODULE=off go build -o build/LibUnitTest_example LibUnitTest_example.go 

echo "Test C++ library"
./build/LibUnitTest_example $PWD/../../Implementations/Cpp/build/libunittest$OSLIBEXT

echo "Test Pascal library"
./build/LibUnitTest_example $PWD/../../Implementations/Pascal/build/libunittest$OSLIBEXT
