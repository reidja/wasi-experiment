This is a very simple experiment using WASI and various guest modules built in Go, Rust, and Zig

There is a go runtime (wazero) that can run all the guest modules

There is a simple `hello-world` style Rust host (wasmer)

There are three examples with various levels of features (go with the most):

    * go-guest-example
    * rust-guest-example
    * zig-guest-example

## Lessons learned

* gRPC was difficult to get working so opted for karmem instead - karmem is pretty obscure would be better to use something more mainstream
* A generator style approach like go-plugin would be more appropriate as its difficult to maintain the interfaces
* WASM only supports ints so you need to move around a lot of bytes to interact between the host and guest
* Error handling is implemented very naively, errors just cause panics

## Requirements

* Rust
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
