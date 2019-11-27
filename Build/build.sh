#!/bin/bash

startingpath="$(pwd)"
basepath="$(cd "$(dirname "$0")" && pwd)"
cd "$basepath/../Source"

Sources="actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go"
GOARCH="amd64"

echo "Build act.exe"
GOOS="windows"
go build -o ../act.exe $Sources

echo "Build act.linux"
GOOS="linux"
go build -o ../act.linux $Sources

echo "Build act.darwin"
GOOS="darwin"
go build -o ../act.darwin $Sources

echo "Build act.arm"
GOOS="linux"
GOARCH="arm"
GOARM="5"
go build -o ../act.arm $Sources

cd "$startingpath"
