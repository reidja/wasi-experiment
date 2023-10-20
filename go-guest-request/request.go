package request

import (
	"github.com/reidja/wasi_demo/guest/core"
	km "github.com/reidja/wasi_demo/shared/request"
)

//go:wasm-module request
//export request_makerequest
func host_request_makerequest(ipos, ilen uint32, rpos **uint32, rlen *uint32) uint32

func guest_request_makerequest(ipos, ilen uint32, rpos **uint32, rlen *uint32) uint32 {
	return host_request_makerequest(ipos, ilen, rpos, rlen)
}

func Get(req *km.Request) km.Response {
	res := km.NewResponse()
	core.SISO(req, &res, guest_request_makerequest)
	return res
}
