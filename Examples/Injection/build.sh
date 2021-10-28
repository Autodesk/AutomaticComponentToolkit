#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT Numbers.xml -suppressexamples
../../$ACT Calculation.xml

echo "Build C++ libraries"
./Numbers_component/Implementations/Cpp/build.sh
./Calculation_component/Implementations/Cpp/build.sh

echo "Build C++ libraries"
./Numbers_component/Implementations/Pascal/build.sh
./Calculation_component/Implementations/Pascal/build.sh


echo "Build and test bindings examples with C++ library"
./Calculation_component/Examples/Cpp/build.sh
./Calculation_component/Examples/CppDynamic/build.sh
./Calculation_component/Examples/Python/build.sh
./Calculation_component/Examples/Pascal/build.sh
./Calculation_component/Examples/Java9/build.sh

echo "Build and test are done and successful"