module github.com/reidja/wasi_demo/guest/cache

go 1.21.3

replace github.com/reidja/wasi_demo/guest/core => ../go-guest-core

require github.com/reidja/wasi_demo/guest/core v0.0.0-00010101000000-000000000000

require karmem.org v1.2.9 // indirect
