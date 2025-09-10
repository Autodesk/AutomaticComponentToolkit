#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build C# example"

[ -d bin ] && rm -rf bin
[ -d obj ] && rm -rf obj
dotnet restore LibPrimes_Example.csproj
dotnet build LibPrimes_Example.csproj --configuration Debug

pushd bin/Debug/net6.0
echo "Test C++ library"
rm -f libprimes.dll
ln -s ../../../../../Implementations/Cpp/build/libprimes$OSLIBEXT libprimes.dll
RUN "dotnet LibPrimes_Example.dll" .

echo "Test Pascal library"
rm -f libprimes.dll
ln -s ../../../../../Implementations/Pascal/build/libprimes$OSLIBEXT libprimes.dll
RUN "dotnet LibPrimes_Example.dll" .

popd
