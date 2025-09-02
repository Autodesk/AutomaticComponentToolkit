#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../Build/build.inc

echo "Generate IDL"
../../$ACT threadSafeLibrary.xml

echo "Build libraries"
./LibThreadSafe_component/Implementations/Cpp/build.sh

echo "Build and test bindings examples with C++ library"
./LibThreadSafe_component/Examples/CppDynamic/build.sh
echo "Build and test are done and successful"
