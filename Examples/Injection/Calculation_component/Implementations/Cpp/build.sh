#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"

echo "Build C++ implementation"
[ -d build ] && rm -rf build
mkdir build
cmake -H. -Bbuild -DCMAKE_BUILD_TYPE=Debug -G Ninja
cmake --build build
