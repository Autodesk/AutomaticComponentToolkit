#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT OptionalClass.xml

echo "Build libraries"
./OptClass_component/Implementations/Cpp/build.sh
./OptClass_component/Implementations/Pascal/build.sh

echo "Build and test bindings examples with C++ library"
./OptClass_component/Examples/Cpp/build.sh
./OptClass_component/Examples/CppDynamic/build.sh
./OptClass_component/Examples/Go/build.sh
./OptClass_component/Examples/Pascal/build.sh
./OptClass_component/Examples/Python/build.sh
echo "Build and test are done and successful"
