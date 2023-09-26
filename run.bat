CALL Build/build.bat
act.exe tmp/libPrimes.xml -o tmp
cd tmp/libPrimes_component/Implementations/Rust
cargo +nightly b