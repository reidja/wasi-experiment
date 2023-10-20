package console

import (
	"context"
	"fmt"

	"github.com/reidja/wasi_demo/host/core"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func Println(ctx context.Context, module api.Module, offset uint32, byteCount uint32) {
	buf := core.ReadToBuffer(module, offset, byteCount)
	fmt.Println(string(buf))
}

func InstantiateModule(ctx context.Context, runtime wazero.Runtime) {
	runtime.
		NewHostModuleBuilder("console").
		NewFunctionBuilder().
		WithFunc(Println).
		Export("console_println").
		Instantiate(ctx)
}
