# Packaging

This folder contains Windows binaries for `libxml2`, the main dependecny of ACT.

Sources:
- http://xmlsoft.org/downloads.html 
- https://www.zlatkovic.com/pub/libxml/64bit/

**Note**

`libxml2-2.dll` has been renamed to `libxml2-2__.dll`, to comply with the naming scheme of the development environment used to build `act.exe`.

## Building the binaries for a release of ACT
1. Build ACT on windows, linux and mac using the scripts in [Build](..)
2. Sign the windows binary (optional)
3. Package the windows binaries and the shared libraries for libxml2 into a self extracting archive (e.g. with WinRar) with the name `ACT_package_win64.exe`.
4. Sign the self extracting package (optional)

