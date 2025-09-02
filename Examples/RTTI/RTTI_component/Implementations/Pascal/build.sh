#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal implementation"
[ -d build ] && rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FuInterfaces -FuStub -FU./build -o./build/rtti$OSLIBEXT Interfaces/rtti.lpr
