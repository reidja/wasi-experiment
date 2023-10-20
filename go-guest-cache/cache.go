package cache

import (
	"github.com/reidja/wasi_demo/guest/core"
)

//go:wasm-module cache
//export cache_get
func host_cache_get(ipos, ilen uint32, rpos **uint32, rlen *uint32) uint32

func guest_cache_get(ipos, ilen uint32, rpos **uint32, rlen *uint32) uint32 {
	return host_cache_get(ipos, ilen, rpos, rlen)
}

func Get(key string) string {
	return core.StrIO(key, guest_cache_get)
}
