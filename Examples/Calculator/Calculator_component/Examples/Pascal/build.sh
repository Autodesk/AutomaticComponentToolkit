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
if [ -f "../../../Implementations/Cpp/build/calculator$OSLIBEXT" ]; then
    rm -f calculator.dll
    ln -s ../../../Implementations/Cpp/build/calculator$OSLIBEXT calculator.dll
    RUN ./Calculator_Example .
else
    echo "C++ library not found - skipping C++ test"
fi

echo "Test Pascal library"
if [ -f "../../../Implementations/Pascal/build/calculator$OSLIBEXT" ]; then
    rm -f calculator.dll
    ln -s ../../../Implementations/Pascal/build/calculator$OSLIBEXT calculator.dll
    RUN ./Calculator_Example .
else
    echo "Pascal library not found - skipping Pascal test (normal if fpc not available)"
fi

popd
