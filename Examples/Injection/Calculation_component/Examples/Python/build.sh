#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Test C++ library"
RUN "python3 Calculation_Example.py" ../../../Numbers_component/Implementations/Cpp/build ../../../Calculation_component/Implementations/Cpp/build 

echo "Test Pascal library"
RUN "python3 Calculation_Example.py" ../../../Numbers_component/Implementations/Pascal/build ../../../Calculation_component/Implementations/Pascal/build
