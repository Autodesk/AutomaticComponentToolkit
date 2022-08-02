#!/bin/bash

failed() {
  echo "${1}" 1>&2
  exit 1;
}

startingpath="$(pwd)"
basepath="$(cd "$(dirname "$0")" && pwd)"
cd "$basepath/../Source"

Sources="actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingccppdocumentation.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildbindingjava.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go"
export GOARCH="amd64"

echo "Build act.exe"
export GOOS="windows"
go build -o ../act.exe $Sources || failed "Error compiling act.exe"

echo "Build act.linux"
export GOOS="linux"
go build -o ../act.linux $Sources || failed "Error compiling act.linux"

echo "Build act.darwin"
export GOOS="darwin"
go build -o ../act.darwin $Sources || failed "Error compiling act.darwin"

echo "Build act.arm.darwin"
export GOOS="darwin"
export GOARCH="arm64"
go build -o ../act.arm.darwin $Sources || failed "Error compiling act.arm.darwin"

echo "Build act.arm.linux" || failed "Error compiling act.arm.linux"
export GOOS="linux"
export GOARCH="arm"
export GOARM="5"
go build -o ../act.arm.linux $Sources

cd "$startingpath"
