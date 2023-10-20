package console

import "github.com/reidja/wasi_demo/guest/core"

//go:wasm-module console
//export console_println
func host_console_println(uint32, uint32)

func Println(message string) {
	ptr, size := core.StringToOffset(message)
	host_console_println(ptr, size)
}
