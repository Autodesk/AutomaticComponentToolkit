@echo off
set startingDir="%CD%"
set basepath="%~dp0"

cd %basepath%\..\Source
set Sources=actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildimplementationcpp.go buildbindingjava.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go
set GOARCH=amd64

set GOOS=windows
echo "Build act.exe"
go build -o ..\act.exe %Sources%

echo "Patching properties of act.exe"
..\build\verpatch ..\act.exe /high /va 1.7.0 /pv "1.7.0-develop" /s copyright "(c) 2018-2019 ACT Developers" /s desc "ACT is a code generator for software components" /s productName "Automatic Component Toolkit"

set GOOS=linux
echo "Build act.linux"
go build -o ..\act.linux %Sources%

set GOOS=darwin
echo "Build act.darwin"
go build -o ..\act.darwin %Sources%

set GOOS=linux
set GOARCH=arm
set GOARM=5
echo "Build act.arm.linux"
go build -o ..\act.arm.linux %Sources%

cd %startingDir%
