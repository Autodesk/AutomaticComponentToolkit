## How to Build

Initialize NuGet packages (run once):
```
msbuild /p:Configuration=Debug /t:Restore RTTI_Example.csproj
```

Allow unsafe code during the build:
```
msbuild /p:Configuration=Debug /p:AllowUnsafeBlocks=true RTTI_Example.csproj
```

## Run:
```
DYLD_LIBRARY_PATH=_path_to_the_lib_ mono bin/Debug/netstandard2.0/RTTI_Example.dll
```

## Super command:
```
../../Build/build.sh && rm -fr RTTI_component/Bindings/Go && ../../act.darwin RTTI.xml &&  msbuild /p:Configuration=Debug /p:AllowUnsafeBlocks=true RTTI_component/Examples/CSharp/RTTI_Example.csproj && DYLD_LIBRARY_PATH=_path_to_rtti_lib_ mono RTTI_component/Examples/CSharp/bin/Debug/netstandard2.0/RTTI_Example.dll
```

NOTE: On Mac original library has `rtti.dylib` name. Rename it or make a symlink `rtti.dll`.
