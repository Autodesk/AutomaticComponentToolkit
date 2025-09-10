#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal implementation"
[ -d build ] && rm -rf build
mkdir build

# Check if we're building on Linux and add -fPIC flag
if [ "$OS" = "Linux" ]; then
  echo "Building for Linux with -fPIC flag"
  fpc -Fu../../Bindings/Pascal -FuInterfaces -FuStub -FU./build -o./build/libprimes$OSLIBEXT -fPIC Interfaces/libprimes.lpr
else
  echo "Building for $OS"
  fpc -Fu../../Bindings/Pascal -FuInterfaces -FuStub -FU./build -o./build/libprimes$OSLIBEXT Interfaces/libprimes.lpr
fi
