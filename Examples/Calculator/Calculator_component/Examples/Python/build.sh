#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Test with available libraries"
if [ -d "$PWD/../../Implementations/Cpp/build" ] && [ -f "$PWD/../../Implementations/Cpp/build/calculator$OSLIBEXT" ]; then
    echo "Testing with C++ library"
    RUN "python3 Calculator_Example.py" $PWD/../../Implementations/Cpp/build
elif [ -d "$PWD/../../Implementations/Pascal/build" ] && [ -f "$PWD/../../Implementations/Pascal/build/calculator$OSLIBEXT" ]; then
    echo "Testing with Pascal library"
    RUN "python3 Calculator_Example.py" $PWD/../../Implementations/Pascal/build
else
    echo "No library found - skipping Python test (this is normal if no implementations were built)"
fi
