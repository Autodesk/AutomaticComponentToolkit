#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal example"
rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FU./build -o./build/RTTI_Example$OSEXEEXT RTTI_Example.lpr

pushd build

echo "Test C++ library"
rm -f rtti.dll
ln -s ../../../Implementations/Cpp/build/rtti$OSLIBEXT rtti.dll
RUN ./RTTI_Example .

echo "Test Pascal library"
rm -f rtti.dll
ln -s ../../../Implementations/Pascal/build/rtti$OSLIBEXT rtti.dll
RUN ./RTTI_Example .

popd
