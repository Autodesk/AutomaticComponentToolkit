#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT Calculator.xml

echo "Build libraries (if cmake is available)"
if command -v cmake >/dev/null 2>&1; then
    ./Calculator_component/Implementations/Cpp/build.sh
else
    echo "Skipping C++ implementation (cmake not found)"
fi

if command -v fpc >/dev/null 2>&1; then
    ./Calculator_component/Implementations/Pascal/build.sh
else
    echo "Skipping Pascal implementation (fpc not found)"
fi

echo "Build and test bindings examples"
if command -v cmake >/dev/null 2>&1; then
    ./Calculator_component/Examples/CppDynamic/build.sh
else
    echo "Skipping C++ Dynamic example (cmake not found)"
fi

if command -v fpc >/dev/null 2>&1; then
    ./Calculator_component/Examples/Pascal/build.sh
else
    echo "Skipping Pascal example (fpc not found)"
fi

if command -v python3 >/dev/null 2>&1; then
    ./Calculator_component/Examples/Python/build.sh
else
    echo "Skipping Python example (python3 not found)"
fi
echo "Build and test are done and successful"
