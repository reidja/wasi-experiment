run_go_example:
    just build_go_example
    cd go-host-runtime && just run ../go-guest-example/main.wasm

run_zig_example:
    just build_zig_example
    cd go-host-runtime && just run ../zig-guest-example/main.wasm

run_rust_example:
    just build_rust_example
    cd go-host-runtime && just run ../rust-guest-example/target/wasm32-wasi/release/rust-demo.wasm

run_rust_host:
    cd rust-host-runtime && just run

build_go_example:
    cd go-guest-example && just build

build_zig_example:
    cd zig-guest-example && just build

build_rust_example:
    cd rust-guest-example && just build
