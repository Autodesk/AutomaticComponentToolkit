#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Go example"

[ -d bin ] && rm -rf bin
[ -d obj ] && rm -rf obj
msbuild /p:Configuration=Debug /t:Restore Calculation_Example.csproj
msbuild /p:Configuration=Debug /p:AllowUnsafeBlocks=true Calculation_Example.csproj

pushd bin/Debug/netstandard2.0 
echo "Test C++ library"
rm -f calculation.dll numbers.dll
ln -s ../../../../../../Calculation_component/Implementations/Cpp/build/calculation$OSLIBEXT calculation.dll
ln -s ../../../../../../Numbers_component/Implementations/Cpp/build/numbers$OSLIBEXT numbers.dll
RUN "mono Calculation_Example.dll" .

echo "Test Pascal library"
rm -f calculation.dll numbers.dll
ln -s ../../../../../../Calculation_component/Implementations/Pascal/build/calculation$OSLIBEXT calculation.dll
ln -s ../../../../../../Numbers_component/Implementations/Pascal/build/numbers$OSLIBEXT numbers.dll
RUN "mono Calculation_Example.dll" .

popd