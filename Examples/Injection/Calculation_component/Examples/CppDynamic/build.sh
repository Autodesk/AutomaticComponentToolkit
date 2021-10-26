#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ implementation"
[ -d build ] && rm -rf build
mkdir build
cmake -H. -Bbuild -DCMAKE_BUILD_TYPE=Debug -G Ninja
cmake --build build

echo "Test C++ library"
RUN ./build/CalculationExample_CPPDynamic ../../../Calculation_component/Implementations/Cpp/build ../../../Numbers_component/Implementations/Cpp/build

echo "Test Pascal library"
RUN ./build/CalculationExample_CPPDynamic ../../../Calculation_component/Implementations/Pascal/build ../../../Numbers_component/Implementations/Pascal/build
