#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")"
source ../../../../../Build/build.inc

echo "Build Go example"

[ -d bin ] && rm -rf bin
[ -d obj ] && rm -rf obj
msbuild /p:Configuration=Debug /t:Restore RTTI_Example.csproj
msbuild /p:Configuration=Debug /p:AllowUnsafeBlocks=true RTTI_Example.csproj

pushd bin/Debug/netstandard2.0 
echo "Test C++ library"
rm -f rtti.dll
ln -s ../../../../../Implementations/Cpp/build/rtti$OSLIBEXT rtti.dll
RUN "mono RTTI_Example.dll" .

echo "Test Pascal library"
rm -f rtti.dll
ln -s ../../../../../Implementations/Pascal/build/rtti$OSLIBEXT rtti.dll
RUN "mono RTTI_Example.dll" .

popd