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
		Sizeof{{$T.TypeName}} = int(unsafe.Sizeof({{$T.GoTypeName}}({{$T.ZeroValue}})))
	{{end}}
)

{{range $T := $.TL -}}
// Bytes to {{$T.TypeName}} Slice force convert.
func BytesTo{{$T.TypeName}}Slice(b []byte) []{{$T.GoTypeName}} {
	{{- if eq $T.Sizeof 1}}
	cvLen := len(b)
	cvCap := cap(b)
	{{- else}}
	cvLen := int((len(b)+Sizeof{{$T.TypeName}}-1)/Sizeof{{$T.TypeName}})
	cvCap := int(cap(b) / Sizeof{{$T.TypeName}})
	if cvLen > cvCap {
		cvLen = cvCap
	}
	{{- end}}
	return *(*[]{{$T.GoTypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len: cvLen,
		Cap: cvCap,
	}))
}

// {{$T.TypeName}} Slice to Bytes force convert.
func {{$T.TypeName}}SliceToBytes(s []{{$T.GoTypeName}}) []byte {
	cvLen := len(s)
	cvCap := cap(s)
	{{if ne $T.Sizeof 1 -}}
	if cvLen >= int(maxInt / Sizeof{{$T.TypeName}}) {
		cvLen = maxInt
	} else {
		cvLen *= Sizeof{{$T.TypeName}}
	}
	if cvCap >= int(maxInt / Sizeof{{$T.TypeName}}) {
		cvCap = maxInt
	} else {
		cvCap *= Sizeof{{$T.TypeName}}
	}
	{{- end}}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len: cvLen,
		Cap: cvCap, 
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
