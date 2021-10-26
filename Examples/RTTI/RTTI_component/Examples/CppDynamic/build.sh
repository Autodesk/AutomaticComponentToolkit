#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ Dynamic example"

[ -d build ] && rm -rf build
mkdir build
pushd build 
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -G Ninja
cmake --build .

echo "Test C++ library"
RUN ./RTTIExample_CPPDynamic ../../../Implementations/Cpp/build

echo "Test Pascal library"
RUN ./RTTIExample_CPPDynamic ../../../Implementations/Pascal/build

popd

