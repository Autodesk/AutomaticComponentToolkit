# Release process for ACT

1. Create Release branch

2. Prepare tutorial etc

3. (optional) Pre-release RC, gather feedback + fix stuff

4. Modify release branch such that all links point to master/Actual.Version.Number

5. build:

	1. Update the properties of the target act.exe by adapting the call to `build/verpatch.exe ...` in `build/build.bat`

	2. `build/build.bat`

	3. Ensure the binary's properties are suitable 
	
	4. sign act.exe
	
6. Merge release branch into master (squash)

7. Create github release

8. Merge Release branch into develop (rebase or merge)

9. Update paths and version in all documents on develop branch



__Notes__:

verpatch.exe is an _amazing_ tool to change resource information in EXEs and DLLs

https://github.com/pavel-a/ddverpatch/releases

Call like this
```
verpatch.exe ..\act.exe /high /va 1.5.0 /pv "1.5.0-RC1+buildnumber-5" /s copyright "(c) 2018-2019 ACT Developers" /s desc "ACT is a code generator for software components" /s productName "Automatic Component Toolkit"
```