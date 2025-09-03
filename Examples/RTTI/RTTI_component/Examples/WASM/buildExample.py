#!/usr/bin/env python3
import sys
import subprocess
from pathlib import Path
import uuid
import shutil

def find_emxx():
    from shutil import which
    emxx = which("em++") or which("emcc")
    if not emxx:
        sys.exit("error: could not find 'em++/emcc' on PATH. Run emsdk_env first.")
    return emxx

def main():
    script_dir = Path(__file__).parent.resolve()          # .../Examples/WASM
    component_root = script_dir.parent.parent             # .../*_component
    if not component_root.name.endswith("_component"):
        sys.exit("fatal: expected *_component/Examples/WASM layout")

    impl_dir = component_root / "Implementations" / "Cpp"
    interfaces_dir = impl_dir / "Interfaces"
    stub_dir = impl_dir / "Stub"
    bindings_wasm_dir = component_root / "Bindings" / "WASM"
    bindings_include_dir = component_root / "Bindings"   # for "Cpp/<Base>_implicit.hpp"

    # find exactly one *_bindings.cpp
    bindings_cpp_list = list(bindings_wasm_dir.glob("*_bindings.cpp"))
    if len(bindings_cpp_list) != 1:
        sys.exit("fatal: expected exactly one *_bindings.cpp in Bindings/WASM/")
    bindings_cpp = bindings_cpp_list[0]

    # derive output name: rtti_bindings.cpp -> rttiBinding.js/.wasm
    wasm_base = bindings_cpp.stem.replace("_bindings", "") + "Binding"

    project_lower = component_root.name.replace("_component", "").lower()
    iface_src = [
        interfaces_dir / f"{project_lower}_interfaceexception.cpp",
        interfaces_dir / f"{project_lower}_interfacewrapper.cpp",
    ]
    stub_srcs = sorted(stub_dir.glob("*.cpp"))

    # sanity
    for p in iface_src + [*stub_srcs, bindings_cpp]:
        if not p.exists():
            sys.exit(f"fatal: missing source: {p}")

    # temp build dir
    build_dir = script_dir / f"build_{uuid.uuid4().hex[:12]}"
    build_dir.mkdir(parents=True, exist_ok=True)

    target_js = build_dir / f"{wasm_base}.js"
    out_js = script_dir / f"{wasm_base}.js"
    out_wasm = script_dir / f"{wasm_base}.wasm"

    emxx = find_emxx()

    compile_flags = [
        "-std=c++17",
        "-O3",
        "-I", str(stub_dir),
        "-I", str(interfaces_dir),
        "-I", str(bindings_include_dir),   # resolves #include "Cpp/<Base>_implicit.hpp"
    ]

    link_flags = [
        "-sWASM=1",
        "-sMODULARIZE=1",
        "-sEXPORT_ES6=1",
        "-sUSE_ES6_IMPORT_META=1",
        "-sENVIRONMENT=node",
        "-sALLOW_MEMORY_GROWTH=1",
        "--bind",
        "-sEXCEPTION_CATCHING_ALLOWED=['*']",
        "-sASSERTIONS=1",
        "--no-entry",
    ]

    cmd = [str(emxx)] + compile_flags + ["-o", str(target_js)]
    cmd += [str(p) for p in stub_srcs + iface_src + [bindings_cpp]]
    cmd += link_flags

    print("Running:\n ", " \\\n    ".join(cmd))
    try:
        subprocess.check_call(cmd, cwd=script_dir)
    except subprocess.CalledProcessError as e:
        print(f"\nBuild failed ({e.returncode}). Keeping {build_dir} for inspection.", file=sys.stderr)
        sys.exit(1)

    # copy artifacts next to example and clean temp
    built_js = target_js
    built_wasm = target_js.with_suffix(".wasm")
    if not built_js.exists() or not built_wasm.exists():
        print(f"\nBuild did not produce expected outputs:\n  {built_js}\n  {built_wasm}", file=sys.stderr)
        sys.exit(1)

    shutil.copy2(built_js, out_js)
    shutil.copy2(built_wasm, out_wasm)

    try:
        shutil.rmtree(build_dir)
    except Exception as _:
        # non-fatal
        pass

    print(f"\n✓ Built {out_js.name} and {out_wasm.name} into {script_dir}")
    print("Run your example with:\n  node example.mjs")

if __name__ == "__main__":
    main()
