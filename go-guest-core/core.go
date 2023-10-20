package core

import (
	"log"
	"unsafe"

	karmem "karmem.org/golang"
)

// convert a string into a "fat" pointer
func StringToOffset(message string) (uint32, uint32) {
	buf := []byte(message)
	return BufferToPtr(buf)
}

// convert a buffer into a fat pointer
func BufferToPtr(buf []byte) (uint32, uint32) {
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

// allocate a buffer within the guest
//
//export AllocateBuffer
func AllocateBuffer(size uint32) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

// read a "fat" pointer from memory into a buffer
func ReadBufferFromMemory(ptr *uint32, size uint32) []byte {
	buf := make([]byte, size)
	pointer := uintptr(unsafe.Pointer(ptr))
	for i := 0; i < int(size); i++ {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(i)))
		buf[i] = byte(s)
	}
	return buf
}

// helper interface to allow request/response using a karmem reader
type Serializable interface {
	ReadAsRoot(*karmem.Reader)
}

// helper interface to allow request/response using a karmem reader
type Deserializable interface {
	WriteAsRoot(*karmem.Writer) (uint, error)
}

// fixed size Karmem reader
var OutputMemory [8_000_000]byte
var _Writer = karmem.NewFixedWriter(OutputMemory[:])

// single input single output using a karmem object as input and writing the response into the buffer
func SISO[A Deserializable, T Serializable](input A, output T, handler func(uint32, uint32, **uint32, *uint32) uint32) {
	// reset the writer and serialize the input
	_Writer.Reset()
	if _, err := input.WriteAsRoot(_Writer); err != nil {
		log.Panicf("unable to write Request")
	}

	// convert the request into a fat pointer
	iptr, isize := BufferToPtr(_Writer.Bytes())

	// locations where the response will be stored
	var rptr *uint32
	var rsize uint32

	// execute the host handler that performs the "request" and writes its response into the fat pointer
	handler(iptr, isize, &rptr, &rsize)

	// read the response pointer and size into a buffer
	buf := ReadBufferFromMemory(rptr, rsize)

	// deserialize the response into output
	reader := karmem.NewReader(buf)
	output.ReadAsRoot(reader)
}

// string input and string output - see SISO
func StrIO(input string, handler func(uint32, uint32, **uint32, *uint32) uint32) string {
	iptr, isize := BufferToPtr([]byte(input))

	var rptr *uint32
	var rsize uint32

	handler(iptr, isize, &rptr, &rsize)

	buf := ReadBufferFromMemory(rptr, rsize)
	return string(buf)
}
