This is an experiment using WASI and various guest modules built in Go, Rust, and Zig

There is a go runtime (wazero) that can run all the guest modules, and a `hello-world` style Rust host (wasmer).

There are three examples with various levels of features (go with the most):
    * go-guest-example
    * rust-guest-example
    * zig-guest-example

## Lessons learned

* gRPC was difficult to get working so opted for Karmem instead
* A generator style approach like go-plugin would be more appropriate as its difficult to maintain the interfaces

## Requirements

* Rust (nightly)
* Just
* Go 1.21
* Tinygo
* Zig

## Usage

You can run all the commands using Just:

```
just --list
    build_go_example
    build_rust_example
    build_zig_example
    run_go_example
    run_rust_example
    run_rust_host
    run_zig_example
```
