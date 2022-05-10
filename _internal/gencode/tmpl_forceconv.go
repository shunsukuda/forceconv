package forceconv

import (
	"reflect"
	"unsafe"
)
const (
	maxInt = int(^uint(0) >> 1)
)

const (
	{{range $T := $.TL -}}
		SizeOf{{$T.TypeName}} = int(unsafe.Sizeof({{$T.GoTypeName}}({{$T.ZeroValue}})))
	{{end}}
)

{{range $T := $.TL -}}
// Bytes to {{$T.TypeName}} Slice force convert.
func BytesTo{{$T.TypeName}}Slice(b []byte) []{{$T.GoTypeName}} {
	{{- if eq $T.SizeOf 1}}
	cvl := len(b)
	cvc := cap(b)
	{{- else}}
	cvl := int(len(b) / SizeOf{{$T.TypeName}})
	cvc := int(cap(b) / SizeOf{{$T.TypeName}})
	{{- end}}
	return *(*[]{{$T.GoTypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len: cvl,
		Cap: cvc,
	}))
}

// {{$T.TypeName}} Slice to Bytes force convert.
func {{$T.TypeName}}SliceToBytes(s []{{$T.GoTypeName}}) []byte {
	cvl := len(s)
	cvc := cap(s)
	{{- if ne $T.SizeOf 1}}
	if cvl >= int(maxInt / SizeOf{{$T.TypeName}}) {
		cvl = maxInt
	} else {
		cvl *= SizeOf{{$T.TypeName}}
	}
	if cvc >= int(maxInt / SizeOf{{$T.TypeName}}) {
		cvc = maxInt
	} else {
		cvc *= SizeOf{{$T.TypeName}}
	}
	{{- end}}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len: cvl,
		Cap: cvc, 
	}))
}

{{end}}

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
