package forceconv

import (
	"testing"
	"bytes"
	"reflect"
)

{{range $T := $.TL -}}
func TestBytesTo{{$T.TypeName}}Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []{{$T.GoTypeName}}
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]{{$T.GoTypeName}}, int(1/SizeOf{{$T.TypeName}})), wantLen: int(1/SizeOf{{$T.TypeName}}), wantCap: int(5/SizeOf{{$T.TypeName}})},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]{{$T.GoTypeName}}, int(30/SizeOf{{$T.TypeName}})), wantLen: int(30/SizeOf{{$T.TypeName}}), wantCap: int(50/SizeOf{{$T.TypeName}})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesTo{{$T.TypeName}}Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTo{{$T.GoTypeName}}Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesTo{{$T.TypeName}}Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesTo{{$T.TypeName}}Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func Test{{$T.TypeName}}SliceToBytes(t *testing.T) {
	type args struct {
		s []{{$T.GoTypeName}}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]{{$T.GoTypeName}}, 1, 5)}, want: make([]byte, 1*SizeOf{{$T.TypeName}}), wantLen: 1*SizeOf{{$T.TypeName}}, wantCap: 5*SizeOf{{$T.TypeName}}},
		{name: "", args: args{s: make([]{{$T.GoTypeName}}, 30, 50)}, want: make([]byte, 30*SizeOf{{$T.TypeName}}), wantLen: 30*SizeOf{{$T.TypeName}}, wantCap: 50*SizeOf{{$T.TypeName}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := {{$T.TypeName}}SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTo{{$T.GoTypeName}}Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesTo{{$T.TypeName}}Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesTo{{$T.TypeName}}Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
{{end}} {{/* range $T */}}
func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{name: "", args: args{b: nil}, want: "", wantLen: 0},
		{name: "", args: args{b: []byte("")}, want: "", wantLen: 0},
		{name: "", args: args{b: []byte("hello")}, want: "hello", wantLen: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToString(tt.args.b)
			l := len(got)
			if got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToString() len = %v, want %v", l, tt.wantLen)
				}
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: ""}, want: []byte(""), wantLen: 0},
		{name: "", args: args{s: "hello"}, want: []byte("hello"), wantLen: 5, wantCap: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToBytes(tt.args.s)
			l := len(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("StringToBytes() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if l != tt.wantCap {
					t.Errorf("StringToBytes() len = %v, want %v", l, tt.wantCap)
				}
			}
		})
	}
}