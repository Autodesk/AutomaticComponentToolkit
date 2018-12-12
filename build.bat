@echo off
cd Source
set Sources=actutils.go automaticcomponenttoolkit.go buildbindingcdynamic.go buildbindingcpp.go buildbindinggo.go buildbindingnode.go buildbindingpascal.go buildbindingpython.go buildimplementationcpp.go buildimplementationpascal.go componentdefinition.go languagewriter.go languagec.go languagepascal.go
set GOOS=windows
set GOARCH=amd64
go build -o ..\act.exe %Sources%
set GOOS=linux
go build -o ..\act.linux %Sources%
set GOOS=darwin
go build -o ..\act.darwin %Sources%
cd ..