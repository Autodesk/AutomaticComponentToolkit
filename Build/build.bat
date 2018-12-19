@echo off
set startingDir="%CD%"
set basepath="%~dp0"

cd %basepath%\..\Source
set Sources=actutils.go^
  automaticcomponenttoolkit.go^
  buildbindingcdynamic.go^
  buildbindingcpp.go^
  buildbindinggo.go^
  buildbindingnode.go^
  buildbindingpascal.go^
  buildbindingpython.go^
  buildimplementationcpp.go^
  buildimplementationpascal.go^
  componentdefinition.go^
  componentdiff.go^
  languagewriter.go^
  languagec.go^
  languagepascal.go^
  schemavalidation.go
set GOARCH=amd64

set LDFLAGS=-ldflags "-s -w"

set GOOS=windows
set CGO_ENABLED=1
echo "Build act.exe"
go build %LDFLAGS% -o ..\act.exe %Sources%

echo GOGCCFLAGS=%GOGCCFLAGS%
set GOGCCFLAGS=-Wno-error -w -fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=C:\Users\weismam\AppData\Local\Temp\go-build679076322=/tmp/go-build -gno-record-gcc-switches -fno-common -w
echo GOGCCFLAGS=%GOGCCFLAGS%

set GOOS=linux
echo "Build act.linux"
go build %LDFLAGS% -o ..\act.linux %Sources%

set GOOS=darwin
echo "Build act.darwin"
go build %LDFLAGS% -o ..\act.darwin %Sources%

cd %startingDir%
