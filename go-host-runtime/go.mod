module github.com/reidja/wasi_demo/host/runtime

go 1.21.3

replace github.com/reidja/wasi_demo/host/console => ../go-host-console

replace github.com/reidja/wasi_demo/host/log => ../go-host-log

replace github.com/reidja/wasi_demo/host/cache => ../go-host-cache

replace github.com/reidja/wasi_demo/host/request => ../go-host-request

require (
	github.com/tetratelabs/wazero v1.5.0
	github.com/reidja/wasi_demo/host/cache v0.0.0-00010101000000-000000000000
	github.com/reidja/wasi_demo/host/console v0.0.0-00010101000000-000000000000
	github.com/reidja/wasi_demo/host/log v0.0.0-00010101000000-000000000000
	github.com/reidja/wasi_demo/host/request v0.0.0-00010101000000-000000000000
)
