package log

import "github.com/reidja/wasi_demo/guest/core"

//go:wasm-module log
//export log_println
func host_log_println(uint32, uint32)

func Println(message string) {
	ptr, size := core.StringToOffset(message)
	host_log_println(ptr, size)
}
