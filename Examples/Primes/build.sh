#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT libPrimes.xml

echo "Build libraries"
./LibPrimes_component/Implementations/Cpp/build.sh
./LibPrimes_component/Implementations/Pascal/build.sh

echo "Build and test bindings examples with C++ library"
./LibPrimes_component/Examples/CDynamic/build.sh
./LibPrimes_component/Examples/Cpp/build.sh
./LibPrimes_component/Examples/CppDynamic/build.sh
./LibPrimes_component/Examples/CSharp/build.sh
./LibPrimes_component/Examples/Go/build.sh
./LibPrimes_component/Examples/Java9/build.sh
./LibPrimes_component/Examples/Pascal/build.sh
./LibPrimes_component/Examples/Python/build.sh
echo "Build and test are done and successful"
