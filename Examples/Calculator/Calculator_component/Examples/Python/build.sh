#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Test with C++ library"
if [ -f "$PWD/../../Implementations/Cpp/build/calculator$OSLIBEXT" ]; then
    echo "Testing with C++ library"
    RUN "python3 Calculator_Example.py" $PWD/../../Implementations/Cpp/build
else
    echo "C++ library not found - skipping C++ test"
fi

echo "Test with Pascal library"
if [ -f "$PWD/../../Implementations/Pascal/build/calculator$OSLIBEXT" ]; then
    echo "Testing with Pascal library"
    RUN "python3 Calculator_Example.py" $PWD/../../Implementations/Pascal/build
else
    echo "Pascal library not found - skipping Pascal test (normal if fpc not available)"
fi
