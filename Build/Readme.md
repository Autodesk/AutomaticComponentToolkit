# Release process for ACT

1. Create Release branch

2. Prepare tutorial etc

3. (optional) Pre-release RC, gather feedback + fix stuff

4. Modify release branch such that all links point to the master-branch/Actual.Version.Number

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

## Docker

Docker image can be used to build ACT and all its example.

**Build Docker image:**
```
./Build/build.sh
```

**Build ACT and all examples**
```
# go into root folder of this repo
docker run -it --rm -v $PWD:/data act-build
```

**Run Docker image interactively:**
```
# go into root folder of this repo
docker run -it --rm -v $PWD:/data --entrypoint bash act-build -l
```
Source code is available in `/data` directory.

**Useful scripts**
| Script | Description|
|--------|------------|
| `./Build/docker.sh build`    | build docker image |
| `./Build/docker.sh act`      | build ACT binaries |
| `./Build/docker.sh examples` | build and run projects in Examples folder |
| `./Build/docker.sh all`      | build ACT binaries and then build and run projects in Examples folder" |
| `./Build/docker.sh cli`      | open bash session inside Docker with source code mapped to '/data' directory" |
| Command to run in Docker cli mode (sources are mapped to `/data` directory): |
| `./Build/build.sh` | build ACT binaries |
| `./Examples/build.sh` | build and run projects in Examples folder (including updating Bindings and Interfaces) |
| `./Examples/RTTI/build.sh` | build and run projects in Examples/RTTI folder (including updating Bindings and Interfaces) |
| `./Examples/Injection/build.sh` | build and run projects in Examples/RTTI folder (including updating Bindings and Interfaces) |
| `./Examples/RTTI/RTTI_component/Implementations/Cpp/build.sh` | build RTTI C++ library implementation |
| `./Examples/RTTI/RTTI_component/Implementations/Pascal/build.sh` | build RTTI Pascal library implementation |
| `./Examples/RTTI/RTTI_component/Examples/CppDynamic/build.sh` | build and run RTTI C++ Example (requres C++ and Pascal libraries) |
| `./Examples/RTTI/RTTI_component/Examples/Python/build.sh` | build and run RTTI Python Example (requres C++ and Pascal libraries) |
