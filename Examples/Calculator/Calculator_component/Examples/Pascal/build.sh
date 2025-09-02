#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal example"
rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FU./build -o./build/Calculator_Example$OSEXEEXT Calculator_Example.lpr

pushd build

echo "Test C++ library"
rm -f calculator.dll
ln -s ../../../Implementations/Cpp/build/calculator$OSLIBEXT calculator.dll
RUN ./Calculator_Example .

echo "Test Pascal library"
rm -f calculator.dll
ln -s ../../../Implementations/Pascal/build/calculator$OSLIBEXT calculator.dll
RUN ./Calculator_Example .

popd
