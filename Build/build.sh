#!/bin/bash

function failed {
  echo "$1" 1>&2
  exit 1;
}

startingpath="$(pwd)"
basepath="$(cd "$(dirname "$0")" && pwd)"
cd "$basepath/../Source"

Sources="actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildbindingjava.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go"
GOARCH="amd64"

echo "Build act.exe"
GOOS="windows"
go build -o ../act.exe $Sources || failed "Error compiling act.exe"

echo "Build act.linux"
GOOS="linux"
go build -o ../act.linux $Sources || failed "Error compiling act.linux"

echo "Build act.darwin"
GOOS="darwin"
go build -o ../act.darwin $Sources || failed "Error compiling act.darwin"

echo "Build act.arm" || failed "Error compiling act.arm"
GOOS="linux"
GOARCH="arm"
GOARM="5"
go build -o ../act.arm $Sources

cd "$startingpath"
