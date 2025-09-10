#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ Implicit example"
[ -d build-cpp ] && rm -rf build-cpp
mkdir build-cpp
pushd build-cpp
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DLIBPRIMES_LIB_LOCATION=../../../Implementations/Cpp/build/libprimes$OSLIBEXT
cmake --build .

echo "Test C++ library"
./LibPrimesExample_CPPImplicit
popd


echo "Build C++ Implicit example"
[ -d build-pascal ] && rm -rf build-pascal
mkdir build-pascal
pushd build-pascal
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DLIBPRIMES_LIB_LOCATION=../../../Implementations/Pascal/build/libprimes$OSLIBEXT
cmake --build .

echo "Test Pascal library"
./LibPrimesExample_CPPImplicit
popd
