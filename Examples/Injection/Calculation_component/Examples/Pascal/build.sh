#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Pascal example"
rm -rf build
mkdir build
fpc -Fu../../../Calculation_component/Bindings/Pascal -Fu../../../Numbers_component/Bindings/Pascal -fPIC -T$FPC_TARGET -FU./build -o./build/Calculation_Example$OSEXEEXT Calculation_Example.lpr

pushd build

echo "Test C++ library"
rm -f *.dll
ln -s ../../../../Calculation_component/Implementations/Cpp/build/calculation$OSLIBEXT calculation.dll
ln -s ../../../../Numbers_component/Implementations/Cpp/build/numbers$OSLIBEXT numbers.dll

RUN ./Calculation_Example .

echo "Test Pascal library"
rm -f *.dll
ln -s ../../../../Calculation_component/Implementations/Pascal/build/calculation$OSLIBEXT calculation.dll
ln -s ../../../../Numbers_component/Implementations/Pascal/build/numbers$OSLIBEXT numbers.dll

RUN ./Calculation_Example .

popd