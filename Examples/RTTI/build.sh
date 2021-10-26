#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT RTTI.xml

echo "Build libraries"
./RTTI_component/Implementations/Cpp/build.sh
./RTTI_component/Implementations/Pascal/build.sh

echo "Build and test bindings examples with C++ library"
./RTTI_component/Examples/CDynamic/build.sh
./RTTI_component/Examples/Cpp/build.sh
./RTTI_component/Examples/CppDynamic/build.sh
./RTTI_component/Examples/CSharp/build.sh
./RTTI_component/Examples/Go/build.sh
./RTTI_component/Examples/Java9/build.sh
./RTTI_component/Examples/Pascal/build.sh
./RTTI_component/Examples/Python/build.sh
echo "Build and test are done and successful"