#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ example"

[ -d build ] && rm -rf build
mkdir build
pushd build 
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug
cmake --build .

echo "Test C++ library"
RUN ./LibUnitTestExample_CPPImplicit ../../../Implementations/Cpp/build

echo "Test Pascal library"
RUN ./LibUnitTestExample_CPPImplicit ../../../Implementations/Pascal/build

popd
