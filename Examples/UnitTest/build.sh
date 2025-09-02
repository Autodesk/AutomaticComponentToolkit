#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT libUnitTest.xml

echo "Build libraries"
./LibUnitTest_component/Implementations/Cpp/build.sh
./LibUnitTest_component/Implementations/Pascal/build.sh

echo "Build and test bindings examples with C++ library"
./LibUnitTest_component/Examples/CDynamic/build.sh
./LibUnitTest_component/Examples/Cpp/build.sh
./LibUnitTest_component/Examples/CppDynamic/build.sh
./LibUnitTest_component/Examples/Go/build.sh
./LibUnitTest_component/Examples/Pascal/build.sh
./LibUnitTest_component/Examples/Python/build.sh
echo "Build and test are done and successful"
