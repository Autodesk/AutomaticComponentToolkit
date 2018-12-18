#!/bin/bash

startingpath="$(pwd)"
basepath="$(cd "$(dirname "$0")" && pwd)"
cd "$basepath/../Source"

Sources="actutils.go automaticcomponenttoolkit.go buildbindingcdynamic.go buildbindingcpp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagepascal.go schemavalidation.go"
GOARCH="amd64"
LDFLAGS1=-ldflags
LDFLAGS2="-s -w"
CGO_ENABLED=1

echo "Build act.exe"
GOOS="windows"
go build $LDFLAGS1 "$LDFLAGS2" -o ../act.exe $Sources || exit 1

echo "Build act.linux"
GOOS="linux"
go build $LDFLAGS1 "$LDFLAGS2" -o ../act.linux $Sources || exit 1

echo "Build act.darwin"
GOOS="darwin"
go build $LDFLAGS1 "$LDFLAGS2" -o ../act.darwin $Sources || exit 1

cd "$startingpath"
ls -ahl ..
