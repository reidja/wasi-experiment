module github.com/reidja/wasi_demo/host/console

go 1.21.3

replace github.com/reidja/wasi_demo/host/core => ../go-host-core

require (
	github.com/reidja/wasi_demo/host/core v0.0.0-00010101000000-000000000000
	github.com/tetratelabs/wazero v1.5.0
)

require karmem.org v1.2.9 // indirect
