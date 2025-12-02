#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal example"
rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FU./build -o./build/LibUnitTest_Example$OSEXEEXT LibUnitTest_Example.lpr

pushd build

echo "Test C++ library"
rm -f libunittest.dll
ln -s ../../../Implementations/Cpp/build/libunittest$OSLIBEXT libunittest.dll
RUN ./LibUnitTest_Example .

echo "Test Pascal library"
rm -f libunittest.dll
ln -s ../../../Implementations/Pascal/build/libunittest$OSLIBEXT libunittest.dll
RUN ./LibUnitTest_Example .

popd
