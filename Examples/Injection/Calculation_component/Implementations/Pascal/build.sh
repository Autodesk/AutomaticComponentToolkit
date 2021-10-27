#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal implementation"
[ -d build ] && rm -rf build
mkdir build
fpc -Fu../../Bindings/Pascal -Fu../../../Numbers_component/Bindings/Pascal -FuInterfaces -FuStub -fPIC -T$FPC_TARGET -FU./build -o./build/calculation$OSLIBEXT Interfaces/calculation.lpr
echo $OS
