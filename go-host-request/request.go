package request

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/reidja/wasi_demo/host/core"
	km "github.com/reidja/wasi_demo/shared/request"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func handleRequest(request *km.Request) *km.Response {
	resp, err := http.Get(request.Url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return &km.Response{
		Url:        request.Url,
		Method:     km.RequestMethodGet,
		Body:       string(body),
		StatusCode: uint8(resp.StatusCode),
	}
}

var MakeRequest = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	ibuf := core.PtrToBuffer(module, stack[0], stack[1])
	k := km.NewRequest()
	obuf := core.SISO(&k, handleRequest, ibuf)
	core.AllocateAndWriteBuffer(ctx, module, obuf, stack[2], stack[3])
	stack[0] = 0
})

func InstantiateModule(ctx context.Context, runtime wazero.Runtime) {
	runtime.
		NewHostModuleBuilder("request").
		NewFunctionBuilder().
		WithGoModuleFunction(MakeRequest,
			[]api.ValueType{
				api.ValueTypeI32,
				api.ValueTypeI32,
				api.ValueTypeI32,
				api.ValueTypeI32,
			},
			[]api.ValueType{api.ValueTypeI32}).
		Export("request_makerequest").
		Instantiate(ctx)
}
