package forceconv

import (
	"reflect"
	"unsafe"
)

const (
	maxInt = int(^uint(0) >> 1)
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
	SizeofString     = int(unsafe.Sizeof(string("")))
)

// Bytes to Bool Slice force convert.
func BytesToBoolSlice(b []byte) []bool {
	cvLen := len(b)
	cvCap := cap(b)
	return *(*[]bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bool Slice to Bytes force convert.
func BoolSliceToBytes(s []bool) []byte {
	cvLen := len(s)
	cvCap := cap(s)

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Int8 Slice force convert.
func BytesToInt8Slice(b []byte) []int8 {
	cvLen := len(b)
	cvCap := cap(b)
	return *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Int8 Slice to Bytes force convert.
func Int8SliceToBytes(s []int8) []byte {
	cvLen := len(s)
	cvCap := cap(s)

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Int16 Slice force convert.
func BytesToInt16Slice(b []byte) []int16 {
	cvLen := int((len(b) + SizeofInt16 - 1) / SizeofInt16)
	cvCap := int(cap(b) / SizeofInt16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Int16 Slice to Bytes force convert.
func Int16SliceToBytes(s []int16) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofInt16) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofInt16
	}
	if cvCap >= int(maxInt/SizeofInt16) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofInt16
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Int32 Slice force convert.
func BytesToInt32Slice(b []byte) []int32 {
	cvLen := int((len(b) + SizeofInt32 - 1) / SizeofInt32)
	cvCap := int(cap(b) / SizeofInt32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Int32 Slice to Bytes force convert.
func Int32SliceToBytes(s []int32) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofInt32) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofInt32
	}
	if cvCap >= int(maxInt/SizeofInt32) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofInt32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Int64 Slice force convert.
func BytesToInt64Slice(b []byte) []int64 {
	cvLen := int((len(b) + SizeofInt64 - 1) / SizeofInt64)
	cvCap := int(cap(b) / SizeofInt64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Int64 Slice to Bytes force convert.
func Int64SliceToBytes(s []int64) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofInt64) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofInt64
	}
	if cvCap >= int(maxInt/SizeofInt64) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofInt64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Uint8 Slice force convert.
func BytesToUint8Slice(b []byte) []uint8 {
	cvLen := len(b)
	cvCap := cap(b)
	return *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Uint8 Slice to Bytes force convert.
func Uint8SliceToBytes(s []uint8) []byte {
	cvLen := len(s)
	cvCap := cap(s)

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Uint16 Slice force convert.
func BytesToUint16Slice(b []byte) []uint16 {
	cvLen := int((len(b) + SizeofUint16 - 1) / SizeofUint16)
	cvCap := int(cap(b) / SizeofUint16)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Uint16 Slice to Bytes force convert.
func Uint16SliceToBytes(s []uint16) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofUint16) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofUint16
	}
	if cvCap >= int(maxInt/SizeofUint16) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofUint16
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Uint32 Slice force convert.
func BytesToUint32Slice(b []byte) []uint32 {
	cvLen := int((len(b) + SizeofUint32 - 1) / SizeofUint32)
	cvCap := int(cap(b) / SizeofUint32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Uint32 Slice to Bytes force convert.
func Uint32SliceToBytes(s []uint32) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofUint32) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofUint32
	}
	if cvCap >= int(maxInt/SizeofUint32) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofUint32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Uint64 Slice force convert.
func BytesToUint64Slice(b []byte) []uint64 {
	cvLen := int((len(b) + SizeofUint64 - 1) / SizeofUint64)
	cvCap := int(cap(b) / SizeofUint64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Uint64 Slice to Bytes force convert.
func Uint64SliceToBytes(s []uint64) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofUint64) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofUint64
	}
	if cvCap >= int(maxInt/SizeofUint64) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofUint64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Float32 Slice force convert.
func BytesToFloat32Slice(b []byte) []float32 {
	cvLen := int((len(b) + SizeofFloat32 - 1) / SizeofFloat32)
	cvCap := int(cap(b) / SizeofFloat32)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Float32 Slice to Bytes force convert.
func Float32SliceToBytes(s []float32) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofFloat32) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofFloat32
	}
	if cvCap >= int(maxInt/SizeofFloat32) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofFloat32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Float64 Slice force convert.
func BytesToFloat64Slice(b []byte) []float64 {
	cvLen := int((len(b) + SizeofFloat64 - 1) / SizeofFloat64)
	cvCap := int(cap(b) / SizeofFloat64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Float64 Slice to Bytes force convert.
func Float64SliceToBytes(s []float64) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofFloat64) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofFloat64
	}
	if cvCap >= int(maxInt/SizeofFloat64) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofFloat64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Complex64 Slice force convert.
func BytesToComplex64Slice(b []byte) []complex64 {
	cvLen := int((len(b) + SizeofComplex64 - 1) / SizeofComplex64)
	cvCap := int(cap(b) / SizeofComplex64)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Complex64 Slice to Bytes force convert.
func Complex64SliceToBytes(s []complex64) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofComplex64) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofComplex64
	}
	if cvCap >= int(maxInt/SizeofComplex64) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofComplex64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to Complex128 Slice force convert.
func BytesToComplex128Slice(b []byte) []complex128 {
	cvLen := int((len(b) + SizeofComplex128 - 1) / SizeofComplex128)
	cvCap := int(cap(b) / SizeofComplex128)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Complex128 Slice to Bytes force convert.
func Complex128SliceToBytes(s []complex128) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofComplex128) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofComplex128
	}
	if cvCap >= int(maxInt/SizeofComplex128) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofComplex128
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// Bytes to String Slice force convert.
func BytesToStringSlice(b []byte) []string {
	cvLen := int((len(b) + SizeofString - 1) / SizeofString)
	cvCap := int(cap(b) / SizeofString)
	if cvLen > cvCap {
		cvLen = cvCap
	}
	return *(*[]string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvLen,
		Cap:  cvCap,
	}))
}

// String Slice to Bytes force convert.
func StringSliceToBytes(s []string) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	if cvLen >= int(maxInt/SizeofString) {
		cvLen = maxInt
	} else {
		cvLen *= SizeofString
	}
	if cvCap >= int(maxInt/SizeofString) {
		cvCap = maxInt
	} else {
		cvCap *= SizeofString
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvLen,
		Cap:  cvCap,
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
