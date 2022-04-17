package forceconv

import (
	"reflect"
	"unsafe"
)


const (
	{{range $T := $.TL -}}
		Sizeof{{$T.TypeName}} = int(unsafe.Sizeof({{$T.GoTypeName}}({{$T.ZeroValue}})))
	{{end}}
)

{{range $T := $.TL -}}
// Bytes to {{$T.TypeName}} Slice force convert.
func BytesTo{{$T.TypeName}}(b []byte) []{{$T.GoTypeName}} {
	if b == nil {
		return nil
	}
	{{if ne $T.Sizeof 1}}
	cvLen := int((len(b)+Sizeof{{$T.TypeName}}-1)/Sizeof{{$T.TypeName}})
	cvCap := int(cap(b) / Sizeof{{$T.TypeName}})
	if cvLen > cvCap {
		cvLen = cvCap
	}
	b = b[:cvLen:cvCap*Sizeof{{$T.TypeName}}]
	{{- end}}
	return *(*[]{{$T.GoTypeName}})(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
		Len:  len(b),
		Cap:  cap(b),
	}))
}

// {{$T.TypeName}} Slice to Bytes force convert.
func {{$T.TypeName}}sToBytes(s []{{$T.GoTypeName}}) []byte {
	if s == nil {
		return nil
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s){{if ne $T.Sizeof 1}}*Sizeof{{$T.TypeName}}{{end}},
		Cap:  cap(s){{if ne $T.Sizeof 1}}*Sizeof{{$T.TypeName}}{{end}},
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