#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C++ Implicit example"
[ -d build-cpp ] && rm -rf build-cpp
mkdir build-cpp
pushd build-cpp
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DCALCULATIONLOCATION=../../../../Calculation_component/Implementations/Cpp/build/calculation$OSLIBEXT -G Ninja
cmake --build .

echo "Test C++ library"
RUN ./CalculationExample_CPPImplicit ../../../../Numbers_component/Implementations/Cpp/build
popd


echo "Build C++ Implicit example"
[ -d build-pascal ] && rm -rf build-pascal
mkdir build-pascal
pushd build-pascal
cmake -H.. -B. -DCMAKE_BUILD_TYPE=Debug -DCALCULATIONLOCATION=../../../../Calculation_component/Implementations/Pascal/build/calculation$OSLIBEXT -G Ninja
cmake --build .

echo "Test Pascal library"
RUN ./CalculationExample_CPPImplicit ../../../../Numbers_component/Implementations/Pascal/build
popd
