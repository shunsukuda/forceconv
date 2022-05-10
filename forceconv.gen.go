package forceconv

import (
	"reflect"
	"unsafe"
)

const (
	maxInt = int(^uint(0) >> 1)
)

const (
	SizeOfBool       = int(unsafe.Sizeof(bool(false)))
	SizeOfInt8       = int(unsafe.Sizeof(int8(0)))
	SizeOfInt16      = int(unsafe.Sizeof(int16(0)))
	SizeOfInt32      = int(unsafe.Sizeof(int32(0)))
	SizeOfInt64      = int(unsafe.Sizeof(int64(0)))
	SizeOfUint8      = int(unsafe.Sizeof(uint8(0)))
	SizeOfUint16     = int(unsafe.Sizeof(uint16(0)))
	SizeOfUint32     = int(unsafe.Sizeof(uint32(0)))
	SizeOfUint64     = int(unsafe.Sizeof(uint64(0)))
	SizeOfFloat32    = int(unsafe.Sizeof(float32(0)))
	SizeOfFloat64    = int(unsafe.Sizeof(float64(0)))
	SizeOfComplex64  = int(unsafe.Sizeof(complex64(0)))
	SizeOfComplex128 = int(unsafe.Sizeof(complex128(0)))
	SizeOfString     = int(unsafe.Sizeof(string("")))
)

// Bytes to Bool Slice force convert.
func BytesToBoolSlice(b []byte) []bool {
	cvl := len(b)
	cvc := cap(b)
	return *(*[]bool)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bool Slice to Bytes force convert.
func BoolSliceToBytes(s []bool) []byte {
	cvl := len(s)
	cvc := cap(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Int8 Slice force convert.
func BytesToInt8Slice(b []byte) []int8 {
	cvl := len(b)
	cvc := cap(b)
	return *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Int8 Slice to Bytes force convert.
func Int8SliceToBytes(s []int8) []byte {
	cvl := len(s)
	cvc := cap(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Int16 Slice force convert.
func BytesToInt16Slice(b []byte) []int16 {
	cvl := int(len(b) / SizeOfInt16)
	cvc := int(cap(b) / SizeOfInt16)
	return *(*[]int16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Int16 Slice to Bytes force convert.
func Int16SliceToBytes(s []int16) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfInt16) {
		cvl = maxInt
	} else {
		cvl *= SizeOfInt16
	}
	if cvc >= int(maxInt/SizeOfInt16) {
		cvc = maxInt
	} else {
		cvc *= SizeOfInt16
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Int32 Slice force convert.
func BytesToInt32Slice(b []byte) []int32 {
	cvl := int(len(b) / SizeOfInt32)
	cvc := int(cap(b) / SizeOfInt32)
	return *(*[]int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Int32 Slice to Bytes force convert.
func Int32SliceToBytes(s []int32) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfInt32) {
		cvl = maxInt
	} else {
		cvl *= SizeOfInt32
	}
	if cvc >= int(maxInt/SizeOfInt32) {
		cvc = maxInt
	} else {
		cvc *= SizeOfInt32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Int64 Slice force convert.
func BytesToInt64Slice(b []byte) []int64 {
	cvl := int(len(b) / SizeOfInt64)
	cvc := int(cap(b) / SizeOfInt64)
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Int64 Slice to Bytes force convert.
func Int64SliceToBytes(s []int64) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfInt64) {
		cvl = maxInt
	} else {
		cvl *= SizeOfInt64
	}
	if cvc >= int(maxInt/SizeOfInt64) {
		cvc = maxInt
	} else {
		cvc *= SizeOfInt64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Uint8 Slice force convert.
func BytesToUint8Slice(b []byte) []uint8 {
	cvl := len(b)
	cvc := cap(b)
	return *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Uint8 Slice to Bytes force convert.
func Uint8SliceToBytes(s []uint8) []byte {
	cvl := len(s)
	cvc := cap(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Uint16 Slice force convert.
func BytesToUint16Slice(b []byte) []uint16 {
	cvl := int(len(b) / SizeOfUint16)
	cvc := int(cap(b) / SizeOfUint16)
	return *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Uint16 Slice to Bytes force convert.
func Uint16SliceToBytes(s []uint16) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfUint16) {
		cvl = maxInt
	} else {
		cvl *= SizeOfUint16
	}
	if cvc >= int(maxInt/SizeOfUint16) {
		cvc = maxInt
	} else {
		cvc *= SizeOfUint16
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Uint32 Slice force convert.
func BytesToUint32Slice(b []byte) []uint32 {
	cvl := int(len(b) / SizeOfUint32)
	cvc := int(cap(b) / SizeOfUint32)
	return *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Uint32 Slice to Bytes force convert.
func Uint32SliceToBytes(s []uint32) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfUint32) {
		cvl = maxInt
	} else {
		cvl *= SizeOfUint32
	}
	if cvc >= int(maxInt/SizeOfUint32) {
		cvc = maxInt
	} else {
		cvc *= SizeOfUint32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Uint64 Slice force convert.
func BytesToUint64Slice(b []byte) []uint64 {
	cvl := int(len(b) / SizeOfUint64)
	cvc := int(cap(b) / SizeOfUint64)
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Uint64 Slice to Bytes force convert.
func Uint64SliceToBytes(s []uint64) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfUint64) {
		cvl = maxInt
	} else {
		cvl *= SizeOfUint64
	}
	if cvc >= int(maxInt/SizeOfUint64) {
		cvc = maxInt
	} else {
		cvc *= SizeOfUint64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Float32 Slice force convert.
func BytesToFloat32Slice(b []byte) []float32 {
	cvl := int(len(b) / SizeOfFloat32)
	cvc := int(cap(b) / SizeOfFloat32)
	return *(*[]float32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Float32 Slice to Bytes force convert.
func Float32SliceToBytes(s []float32) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfFloat32) {
		cvl = maxInt
	} else {
		cvl *= SizeOfFloat32
	}
	if cvc >= int(maxInt/SizeOfFloat32) {
		cvc = maxInt
	} else {
		cvc *= SizeOfFloat32
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Float64 Slice force convert.
func BytesToFloat64Slice(b []byte) []float64 {
	cvl := int(len(b) / SizeOfFloat64)
	cvc := int(cap(b) / SizeOfFloat64)
	return *(*[]float64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Float64 Slice to Bytes force convert.
func Float64SliceToBytes(s []float64) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfFloat64) {
		cvl = maxInt
	} else {
		cvl *= SizeOfFloat64
	}
	if cvc >= int(maxInt/SizeOfFloat64) {
		cvc = maxInt
	} else {
		cvc *= SizeOfFloat64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Complex64 Slice force convert.
func BytesToComplex64Slice(b []byte) []complex64 {
	cvl := int(len(b) / SizeOfComplex64)
	cvc := int(cap(b) / SizeOfComplex64)
	return *(*[]complex64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Complex64 Slice to Bytes force convert.
func Complex64SliceToBytes(s []complex64) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfComplex64) {
		cvl = maxInt
	} else {
		cvl *= SizeOfComplex64
	}
	if cvc >= int(maxInt/SizeOfComplex64) {
		cvc = maxInt
	} else {
		cvc *= SizeOfComplex64
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to Complex128 Slice force convert.
func BytesToComplex128Slice(b []byte) []complex128 {
	cvl := int(len(b) / SizeOfComplex128)
	cvc := int(cap(b) / SizeOfComplex128)
	return *(*[]complex128)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Complex128 Slice to Bytes force convert.
func Complex128SliceToBytes(s []complex128) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfComplex128) {
		cvl = maxInt
	} else {
		cvl *= SizeOfComplex128
	}
	if cvc >= int(maxInt/SizeOfComplex128) {
		cvc = maxInt
	} else {
		cvc *= SizeOfComplex128
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// Bytes to String Slice force convert.
func BytesToStringSlice(b []byte) []string {
	cvl := int(len(b) / SizeOfString)
	cvc := int(cap(b) / SizeOfString)
	return *(*[]string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  cvl,
		Cap:  cvc,
	}))
}

// String Slice to Bytes force convert.
func StringSliceToBytes(s []string) []byte {
	cvl := len(s)
	cvc := cap(s)
	if cvl >= int(maxInt/SizeOfString) {
		cvl = maxInt
	} else {
		cvl *= SizeOfString
	}
	if cvc >= int(maxInt/SizeOfString) {
		cvc = maxInt
	} else {
		cvc *= SizeOfString
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  cvl,
		Cap:  cvc,
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
