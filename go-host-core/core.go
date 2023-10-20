package core

import (
	"context"
	"log"
	"unsafe"

	"github.com/tetratelabs/wazero/api"
	karmem "karmem.org/golang"
)

func ReadToBuffer(module api.Module, offset uint32, byteCount uint32) []byte {
	buf, ok := module.Memory().Read(offset, byteCount)
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	return buf
}

func PtrToBuffer(module api.Module, offset uint64, size uint64) []byte {
	rpos := uint32(offset)
	rsize := uint32(size)
	buf, ok := module.Memory().Read(rpos, rsize)
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", rpos, rsize)
	}
	return buf
}

func BufferToOffset(buf []byte) (uint32, uint32) {
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

func AllocateAndWriteBuffer(ctx context.Context, module api.Module, buf []byte, rpos uint64, rlen uint64) {
	bsize := len(buf)
	rbuf, err := module.ExportedFunction("AllocateBuffer").Call(ctx, uint64(bsize))
	if err != nil {
		log.Panicf("unable to allocate buffer: %v", err)
	}

	apos := uint32(rbuf[0])
	roffset := uint32(rpos)
	rsize := uint32(rlen)
	module.Memory().WriteUint32Le(uint32(roffset), apos)
	module.Memory().WriteUint32Le(uint32(rsize), uint32(bsize))
	module.Memory().Write(apos, buf)
}

type Serializable interface {
	ReadAsRoot(*karmem.Reader)
}

type Deserializable interface {
	WriteAsRoot(*karmem.Writer) (uint, error)
}

var OutputMemory [8_000_000]byte
var _Writer = karmem.NewFixedWriter(OutputMemory[:])

func SISO[A Serializable, T Deserializable](input A, handler func(A) T, buf []byte) []byte {
	reader := karmem.NewReader(buf)
	input.ReadAsRoot(reader)
	res := handler(input)
	_Writer.Reset()
	if _, err := res.WriteAsRoot(_Writer); err != nil {
		log.Panicf("unable to write response: %v", err)
	}
	return _Writer.Bytes()
}

func StrIO(input string, handler func(string) string) []byte {
	res := handler(input)
	return []byte(res)
}
