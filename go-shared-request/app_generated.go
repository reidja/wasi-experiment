package request

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = make([]byte, 32)
var _NullReader = karmem.NewReader(_Null)

type (
	RequestMethod uint8
)

const (
	RequestMethodGet  RequestMethod = 0
	RequestMethodPost RequestMethod = 1
)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierRequest  = 14563690110367227615
	PacketIdentifierResponse = 10395374441924781581
)

type Request struct {
	Url string
}

func NewRequest() Request {
	return Request{}
}

func (x *Request) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierRequest
}

func (x *Request) Reset() {
	x.Read((*RequestViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Request) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Request) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(16))
	__UrlSize := uint(1 * len(x.Url))
	__UrlOffset, err := writer.Alloc(__UrlSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__UrlOffset))
	writer.Write4At(offset+4+4, uint32(__UrlSize))
	writer.Write4At(offset+4+4+4, 1)
	__UrlSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Url)), __UrlSize, __UrlSize}
	writer.WriteAt(__UrlOffset, *(*[]byte)(unsafe.Pointer(&__UrlSlice)))

	return offset, nil
}

func (x *Request) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewRequestViewer(reader, 0), reader)
}

func (x *Request) Read(viewer *RequestViewer, reader *karmem.Reader) {
	__UrlString := viewer.Url(reader)
	if x.Url != __UrlString {
		__UrlStringCopy := make([]byte, len(__UrlString))
		copy(__UrlStringCopy, __UrlString)
		x.Url = *(*string)(unsafe.Pointer(&__UrlStringCopy))
	}
}

type Response struct {
	Url        string
	Method     RequestMethod
	Body       string
	StatusCode uint8
}

func NewResponse() Response {
	return Response{}
}

func (x *Response) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierResponse
}

func (x *Response) Reset() {
	x.Read((*ResponseViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Response) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Response) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(30))
	__UrlSize := uint(1 * len(x.Url))
	__UrlOffset, err := writer.Alloc(__UrlSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__UrlOffset))
	writer.Write4At(offset+4+4, uint32(__UrlSize))
	writer.Write4At(offset+4+4+4, 1)
	__UrlSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Url)), __UrlSize, __UrlSize}
	writer.WriteAt(__UrlOffset, *(*[]byte)(unsafe.Pointer(&__UrlSlice)))
	__MethodOffset := offset + 16
	writer.Write1At(__MethodOffset, *(*uint8)(unsafe.Pointer(&x.Method)))
	__BodySize := uint(1 * len(x.Body))
	__BodyOffset, err := writer.Alloc(__BodySize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+17, uint32(__BodyOffset))
	writer.Write4At(offset+17+4, uint32(__BodySize))
	writer.Write4At(offset+17+4+4, 1)
	__BodySlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Body)), __BodySize, __BodySize}
	writer.WriteAt(__BodyOffset, *(*[]byte)(unsafe.Pointer(&__BodySlice)))
	__StatusCodeOffset := offset + 29
	writer.Write1At(__StatusCodeOffset, *(*uint8)(unsafe.Pointer(&x.StatusCode)))

	return offset, nil
}

func (x *Response) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewResponseViewer(reader, 0), reader)
}

func (x *Response) Read(viewer *ResponseViewer, reader *karmem.Reader) {
	__UrlString := viewer.Url(reader)
	if x.Url != __UrlString {
		__UrlStringCopy := make([]byte, len(__UrlString))
		copy(__UrlStringCopy, __UrlString)
		x.Url = *(*string)(unsafe.Pointer(&__UrlStringCopy))
	}
	x.Method = RequestMethod(viewer.Method())
	__BodyString := viewer.Body(reader)
	if x.Body != __BodyString {
		__BodyStringCopy := make([]byte, len(__BodyString))
		copy(__BodyStringCopy, __BodyString)
		x.Body = *(*string)(unsafe.Pointer(&__BodyStringCopy))
	}
	x.StatusCode = viewer.StatusCode()
}

type RequestViewer struct {
	_data [24]byte
}

func NewRequestViewer(reader *karmem.Reader, offset uint32) (v *RequestViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*RequestViewer)(unsafe.Pointer(&_Null))
	}
	v = (*RequestViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*RequestViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *RequestViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *RequestViewer) Url(reader *karmem.Reader) (v string) {
	if 4+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type ResponseViewer struct {
	_data [32]byte
}

func NewResponseViewer(reader *karmem.Reader, offset uint32) (v *ResponseViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*ResponseViewer)(unsafe.Pointer(&_Null))
	}
	v = (*ResponseViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*ResponseViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *ResponseViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *ResponseViewer) Url(reader *karmem.Reader) (v string) {
	if 4+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ResponseViewer) Method() (v RequestMethod) {
	if 16+1 > x.size() {
		return v
	}
	return *(*RequestMethod)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}
func (x *ResponseViewer) Body(reader *karmem.Reader) (v string) {
	if 17+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 17))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 17+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ResponseViewer) StatusCode() (v uint8) {
	if 29+1 > x.size() {
		return v
	}
	return *(*uint8)(unsafe.Add(unsafe.Pointer(&x._data), 29))
}
