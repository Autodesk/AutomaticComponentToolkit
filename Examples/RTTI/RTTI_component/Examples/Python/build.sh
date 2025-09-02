#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Test C++ library"
RUN "python3 RTTI_Example.py" $PWD/../../Implementations/Cpp/build

echo "Test Pascal library"
RUN "python3 RTTI_Example.py" $PWD/../../Implementations/Pascal/build
