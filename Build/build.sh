#!/bin/bash

set -euxo pipefail

startingpath="$(pwd)"
basepath="$(cd "$(dirname "$0")" && pwd)"
cd "$basepath/../Source"

Sources="actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingccppdocumentation.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildbindingjava.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go"

echo "Build act.win64.exe"
export GOARCH="amd64"
export GOOS="windows"
go build -o ../act.win64.exe $Sources

echo "Build act.win32.exe"
export GOARCH="386"
export GOOS="windows"
go build -o ../act.win32.exe $Sources

echo "Build act.linux64"
export GOOS="linux"
export GOARCH="amd64"
go build -o ../act.linux64 $Sources

echo "Build act.linux32"
export GOOS="linux"
export GOARCH="386"
go build -o ../act.linux32 $Sources

echo "Build act.darwin"
export GOOS="darwin"
export GOARCH="amd64"
go build -o ../act.darwin $Sources

echo "Build act.arm.darwin"
export GOOS="darwin"
export GOARCH="arm64"
go build -o ../act.arm.darwin $Sources

echo "Build act.arm.linux32"
export GOOS="linux"
export GOARM="5"
export GOARCH="386"
go build -o ../act.arm.linux32 $Sources

echo "Build act.arm.linux64"
export GOOS="linux"
export GOARCH="arm64"
export GOARM="5"
go build -o ../act.arm.linux64 $Sources

cd "$startingpath"
