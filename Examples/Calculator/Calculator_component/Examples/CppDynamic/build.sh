#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ Dynamic example"

[ -d build ] && rm -rf build
mkdir build
pushd build 
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug
cmake --build .

echo "List C++ library"
ls ../../../Implementations/Cpp/build

echo "Test C++ library"
RUN ./CalculatorExample_CPPDynamic ../../../Implementations/Cpp/build

echo "List Pascal library"
ls ../../../Implementations/Pascal/build

echo "Test Pascal library"
if [ -f "../../../Implementations/Pascal/build/calculator$OSLIBEXT" ]; then
    RUN ./CalculatorExample_CPPDynamic ../../../Implementations/Pascal/build
else
    echo "Pascal library not found - skipping Pascal test (normal if fpc not available)"
fi

popd
