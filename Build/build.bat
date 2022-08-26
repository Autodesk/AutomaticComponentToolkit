@echo off
set startingDir="%CD%"
set basepath="%~dp0"

cd %basepath%\..\Source
set Sources=actutils.go automaticcomponenttoolkit.go buildbindingccpp.go buildbindingccppdocumentation.go buildbindingcsharp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildimplementationcpp.go buildbindingjava.go buildimplementationpascal.go componentdefinition.go componentdiff.go languagewriter.go languagec.go languagecpp.go languagepascal.go

set GOOS=windows
set GOARCH=amd64
echo "Build act.win64.exe"
go build -o ..\act.win64.exe %Sources%

echo "Patching properties of act.win64.exe"
..\build\verpatch ..\act.win64.exe /high /va 1.8.0 /pv "1.8.0-develop" /s copyright "(c) 2018-2019 ACT Developers" /s desc "ACT is a code generator for software components" /s productName "Automatic Component Toolkit"

set GOOS=windows
set GOARCH=386
echo "Build act.win32.exe"
go build -o ..\act.win32.exe %Sources%

echo "Patching properties of act.win32.exe"
..\build\verpatch ..\act.win32.exe /high /va 1.8.0 /pv "1.8.0-develop" /s copyright "(c) 2018-2019 ACT Developers" /s desc "ACT is a code generator for software components" /s productName "Automatic Component Toolkit"

set GOOS=linux
set GOARCH=amd64
echo "Build act.linux64"
go build -o ..\act.linux64 %Sources%

set GOOS=linux
set GOARCH=386
echo "Build act.linux32"
go build -o ..\act.linux32 %Sources%

set GOOS=darwin
set GOARCH=amd64
echo "Build act.darwin"
go build -o ..\act.darwin %Sources%

set GOOS=darwin
set GOARCH=arm
set GOARM=5
echo "Build act.arm.darwin"
go build -o ..\act.arm.darwin %Sources%

set GOOS=linux
set GOARCH=arm
set GOARM=5
go build -o ..\act.arm.linux32 %Sources%

set GOOS=linux
set GOARCH=arm64
set GOARM=5
echo "Build act.arm.linux64"
go build -o ..\act.arm.linux64 %Sources%

copy ..\act.win64.exe ..\act.exe /Y
copy ..\act.linux64 ..\act.linux /Y


cd %startingDir%
