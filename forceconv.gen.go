package forceconv

import (
	"reflect"
	"unsafe"
)

const (
	SizeofBool       = int(unsafe.Sizeof(bool(false)))
	SizeofInt8       = int(unsafe.Sizeof(int8(0)))
	SizeofInt16      = int(unsafe.Sizeof(int16(0)))
	SizeofInt32      = int(unsafe.Sizeof(int32(0)))
	SizeofInt64      = int(unsafe.Sizeof(int64(0)))
	SizeofUint8      = int(unsafe.Sizeof(uint8(0)))
	SizeofUint16     = int(unsafe.Sizeof(uint16(0)))
	SizeofUint32     = int(unsafe.Sizeof(uint32(0)))
	SizeofUint64     = int(unsafe.Sizeof(uint64(0)))
	SizeofFloat32    = int(unsafe.Sizeof(float32(0.0)))
	SizeofFloat64    = int(unsafe.Sizeof(float64(0.0)))
	SizeofComplex64  = int(unsafe.Sizeof(complex64(0)))
	SizeofComplex128 = int(unsafe.Sizeof(complex128(0)))
)

// Bytes to Bool Slice force convert.
func BytesToBoolSlice(b []byte) []bool {
	if b == nil {
		return nil
	}

	return *(*[]bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Bool Slice to Bytes force convert.
func BoolSliceToBytes(s []bool) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

// Bytes to Int8 Slice force convert.
func BytesToInt8Slice(b []byte) []int8 {
	if b == nil {
		return nil
	}

	return *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Int8 Slice to Bytes force convert.
func Int8SliceToBytes(s []int8) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

// Bytes to Int16 Slice force convert.
func BytesToInt16Slice(b []byte) []int16 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt16 - 1) / SizeofInt16)
	cvCap := int(cap(b) / SizeofInt16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt16]
	return *(*[]int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Int16 Slice to Bytes force convert.
func Int16SliceToBytes(s []int16) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt16,
		Cap:  cap(s) * SizeofInt16,
	}))
}

// Bytes to Int32 Slice force convert.
func BytesToInt32Slice(b []byte) []int32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt32 - 1) / SizeofInt32)
	cvCap := int(cap(b) / SizeofInt32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt32]
	return *(*[]int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Int32 Slice to Bytes force convert.
func Int32SliceToBytes(s []int32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt32,
		Cap:  cap(s) * SizeofInt32,
	}))
}

// Bytes to Int64 Slice force convert.
func BytesToInt64Slice(b []byte) []int64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofInt64 - 1) / SizeofInt64)
	cvCap := int(cap(b) / SizeofInt64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofInt64]
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Int64 Slice to Bytes force convert.
func Int64SliceToBytes(s []int64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofInt64,
		Cap:  cap(s) * SizeofInt64,
	}))
}

// Bytes to Uint8 Slice force convert.
func BytesToUint8Slice(b []byte) []uint8 {
	if b == nil {
		return nil
	}

	return *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Uint8 Slice to Bytes force convert.
func Uint8SliceToBytes(s []uint8) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  cap(s),
	}))
}

// Bytes to Uint16 Slice force convert.
func BytesToUint16Slice(b []byte) []uint16 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint16 - 1) / SizeofUint16)
	cvCap := int(cap(b) / SizeofUint16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint16]
	return *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Uint16 Slice to Bytes force convert.
func Uint16SliceToBytes(s []uint16) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint16,
		Cap:  cap(s) * SizeofUint16,
	}))
}

// Bytes to Uint32 Slice force convert.
func BytesToUint32Slice(b []byte) []uint32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint32 - 1) / SizeofUint32)
	cvCap := int(cap(b) / SizeofUint32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint32]
	return *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Uint32 Slice to Bytes force convert.
func Uint32SliceToBytes(s []uint32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint32,
		Cap:  cap(s) * SizeofUint32,
	}))
}

// Bytes to Uint64 Slice force convert.
func BytesToUint64Slice(b []byte) []uint64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofUint64 - 1) / SizeofUint64)
	cvCap := int(cap(b) / SizeofUint64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofUint64]
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Uint64 Slice to Bytes force convert.
func Uint64SliceToBytes(s []uint64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofUint64,
		Cap:  cap(s) * SizeofUint64,
	}))
}

// Bytes to Float32 Slice force convert.
func BytesToFloat32Slice(b []byte) []float32 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofFloat32 - 1) / SizeofFloat32)
	cvCap := int(cap(b) / SizeofFloat32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofFloat32]
	return *(*[]float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Float32 Slice to Bytes force convert.
func Float32SliceToBytes(s []float32) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofFloat32,
		Cap:  cap(s) * SizeofFloat32,
	}))
}

// Bytes to Float64 Slice force convert.
func BytesToFloat64Slice(b []byte) []float64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofFloat64 - 1) / SizeofFloat64)
	cvCap := int(cap(b) / SizeofFloat64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofFloat64]
	return *(*[]float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Float64 Slice to Bytes force convert.
func Float64SliceToBytes(s []float64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofFloat64,
		Cap:  cap(s) * SizeofFloat64,
	}))
}

// Bytes to Complex64 Slice force convert.
func BytesToComplex64Slice(b []byte) []complex64 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofComplex64 - 1) / SizeofComplex64)
	cvCap := int(cap(b) / SizeofComplex64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofComplex64]
	return *(*[]complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Complex64 Slice to Bytes force convert.
func Complex64SliceToBytes(s []complex64) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofComplex64,
		Cap:  cap(s) * SizeofComplex64,
	}))
}

// Bytes to Complex128 Slice force convert.
func BytesToComplex128Slice(b []byte) []complex128 {
	if b == nil {
		return nil
	}

	cvLen := int((len(b) + SizeofComplex128 - 1) / SizeofComplex128)
	cvCap := int(cap(b) / SizeofComplex128)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[: cvLen : cvCap*SizeofComplex128]
	return *(*[]complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// Complex128 Slice to Bytes force convert.
func Complex128SliceToBytes(s []complex128) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s) * SizeofComplex128,
		Cap:  cap(s) * SizeofComplex128,
	}))
}

// Bytes to String force convert.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(
		&reflect.StringHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
			Len:  len(b),
		}))
}

// String to Bytes force convert.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
			Len:  len(s),
			Cap:  len(s),
		}))
}
