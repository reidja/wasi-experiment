package log

import (
	"context"
	"log"

	"github.com/reidja/wasi_demo/host/core"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func Println(ctx context.Context, module api.Module, offset uint32, byteCount uint32) {
	buf := core.ReadToBuffer(module, offset, byteCount)
	log.Println(string(buf))
}

func Count(a uint32, b uint32) uint32 {
	return a + b
}

func InstantiateModule(ctx context.Context, runtime wazero.Runtime) {
	runtime.
		NewHostModuleBuilder("log").
		NewFunctionBuilder().
		WithFunc(Println).
		Export("log_println").
		NewFunctionBuilder().
		WithFunc(Count).
		Export("log_count").
		Instantiate(ctx)
}
