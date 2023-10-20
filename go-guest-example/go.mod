module github.com/reidja/wasi_demo/guest/example

go 1.21.3

replace github.com/reidja/wasi_demo/guest/console => ../go-guest-console
replace github.com/reidja/wasi_demo/guest/log => ../go-guest-log
replace github.com/reidja/wasi_demo/guest/request => ../go-guest-request
replace github.com/reidja/wasi_demo/shared/request => ../go-shared-request
