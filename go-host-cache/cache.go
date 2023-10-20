package log

import (
	"context"

	"github.com/reidja/wasi_demo/host/core"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// Mocked cache
func _get(key string) string {
	return key + ":hello"
}

// Exposed cache.Get(key) -> key method
var Get = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	ibuf := core.PtrToBuffer(module, stack[0], stack[1])
	obuf := core.StrIO(string(ibuf), _get)
	core.AllocateAndWriteBuffer(ctx, module, obuf, stack[2], stack[3])
	stack[0] = 0
})

func InstantiateModule(ctx context.Context, runtime wazero.Runtime) {
	runtime.
		NewHostModuleBuilder("cache").
		NewFunctionBuilder().
		WithGoModuleFunction(Get,
			[]api.ValueType{
				api.ValueTypeI32,
				api.ValueTypeI32,
				api.ValueTypeI32,
				api.ValueTypeI32,
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("cache_get").
		Instantiate(ctx)
}
