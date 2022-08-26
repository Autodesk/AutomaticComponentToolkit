#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ Implicit example"
[ -d build-cpp ] && rm -rf build-cpp
mkdir build-cpp
pushd build-cpp
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DRTTI_LIB_LOCATION=../../../Implementations/Cpp/build/rtti$OSLIBEXT -G Ninja
cmake --build .

echo "Test C++ library"
./RTTIExample_CPPImplicit
popd


echo "Build C++ Implicit example"
[ -d build-pascal ] && rm -rf build-pascal
mkdir build-pascal
pushd build-pascal
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DRTTI_LIB_LOCATION=../../../Implementations/Pascal/build/rtti$OSLIBEXT -G Ninja
cmake --build .

echo "Test C++ library"
./RTTIExample_CPPImplicit
popd
