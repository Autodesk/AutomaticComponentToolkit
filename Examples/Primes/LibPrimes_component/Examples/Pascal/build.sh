#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal example"
rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FU./build -o./build/LibPrimes_Example$OSEXEEXT LibPrimes_Example.lpr

pushd build

echo "Test C++ library"
rm -f libprimes.dll
ln -s ../../../Implementations/Cpp/build/libprimes$OSLIBEXT libprimes.dll
RUN ./LibPrimes_Example .

echo "Test Pascal library"
rm -f libprimes.dll
ln -s ../../../Implementations/Pascal/build/libprimes$OSLIBEXT libprimes.dll
RUN ./LibPrimes_Example .

popd
