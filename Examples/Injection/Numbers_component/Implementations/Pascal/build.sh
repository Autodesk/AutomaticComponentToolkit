#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal implementation"
[ -d build ] && rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -FuInterfaces -FuStub  -fPIC -T$FPC_TARGET -FU./build -o./build/numbers$OSLIBEXT Interfaces/numbers.lpr
