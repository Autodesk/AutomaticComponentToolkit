## How to Build

`GO111MODULE=off` environment variable must be set to allow local modules.

```
DYLD_LIBRARY_PATH=_path_to_the_lib_ GO111MODULE=off go run RTTI_example.go
```