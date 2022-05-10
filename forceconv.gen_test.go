package forceconv

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBytesToBoolSlice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]bool, int(1/SizeOfBool)), wantLen: int(1 / SizeOfBool), wantCap: int(5 / SizeOfBool)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]bool, int(30/SizeOfBool)), wantLen: int(30 / SizeOfBool), wantCap: int(50 / SizeOfBool)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToBoolSlice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToboolSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToBoolSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToBoolSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestBoolSliceToBytes(t *testing.T) {
	type args struct {
		s []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]bool, 1, 5)}, want: make([]byte, 1*SizeOfBool), wantLen: 1 * SizeOfBool, wantCap: 5 * SizeOfBool},
		{name: "", args: args{s: make([]bool, 30, 50)}, want: make([]byte, 30*SizeOfBool), wantLen: 30 * SizeOfBool, wantCap: 50 * SizeOfBool},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolSliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesToboolSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToBoolSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToBoolSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToInt8Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []int8
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]int8, int(1/SizeOfInt8)), wantLen: int(1 / SizeOfInt8), wantCap: int(5 / SizeOfInt8)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]int8, int(30/SizeOfInt8)), wantLen: int(30 / SizeOfInt8), wantCap: int(50 / SizeOfInt8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToInt8Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToint8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt8SliceToBytes(t *testing.T) {
	type args struct {
		s []int8
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]int8, 1, 5)}, want: make([]byte, 1*SizeOfInt8), wantLen: 1 * SizeOfInt8, wantCap: 5 * SizeOfInt8},
		{name: "", args: args{s: make([]int8, 30, 50)}, want: make([]byte, 30*SizeOfInt8), wantLen: 30 * SizeOfInt8, wantCap: 50 * SizeOfInt8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesToint8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToInt16Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []int16
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]int16, int(1/SizeOfInt16)), wantLen: int(1 / SizeOfInt16), wantCap: int(5 / SizeOfInt16)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]int16, int(30/SizeOfInt16)), wantLen: int(30 / SizeOfInt16), wantCap: int(50 / SizeOfInt16)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToInt16Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToint16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt16SliceToBytes(t *testing.T) {
	type args struct {
		s []int16
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]int16, 1, 5)}, want: make([]byte, 1*SizeOfInt16), wantLen: 1 * SizeOfInt16, wantCap: 5 * SizeOfInt16},
		{name: "", args: args{s: make([]int16, 30, 50)}, want: make([]byte, 30*SizeOfInt16), wantLen: 30 * SizeOfInt16, wantCap: 50 * SizeOfInt16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesToint16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToInt32Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]int32, int(1/SizeOfInt32)), wantLen: int(1 / SizeOfInt32), wantCap: int(5 / SizeOfInt32)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]int32, int(30/SizeOfInt32)), wantLen: int(30 / SizeOfInt32), wantCap: int(50 / SizeOfInt32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToInt32Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToint32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt32SliceToBytes(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]int32, 1, 5)}, want: make([]byte, 1*SizeOfInt32), wantLen: 1 * SizeOfInt32, wantCap: 5 * SizeOfInt32},
		{name: "", args: args{s: make([]int32, 30, 50)}, want: make([]byte, 30*SizeOfInt32), wantLen: 30 * SizeOfInt32, wantCap: 50 * SizeOfInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesToint32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToInt64Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]int64, int(1/SizeOfInt64)), wantLen: int(1 / SizeOfInt64), wantCap: int(5 / SizeOfInt64)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]int64, int(30/SizeOfInt64)), wantLen: int(30 / SizeOfInt64), wantCap: int(50 / SizeOfInt64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToInt64Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToint64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestInt64SliceToBytes(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]int64, 1, 5)}, want: make([]byte, 1*SizeOfInt64), wantLen: 1 * SizeOfInt64, wantCap: 5 * SizeOfInt64},
		{name: "", args: args{s: make([]int64, 30, 50)}, want: make([]byte, 30*SizeOfInt64), wantLen: 30 * SizeOfInt64, wantCap: 50 * SizeOfInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesToint64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToInt64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToInt64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToUint8Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []uint8
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]uint8, int(1/SizeOfUint8)), wantLen: int(1 / SizeOfUint8), wantCap: int(5 / SizeOfUint8)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]uint8, int(30/SizeOfUint8)), wantLen: int(30 / SizeOfUint8), wantCap: int(50 / SizeOfUint8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToUint8Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTouint8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint8SliceToBytes(t *testing.T) {
	type args struct {
		s []uint8
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]uint8, 1, 5)}, want: make([]byte, 1*SizeOfUint8), wantLen: 1 * SizeOfUint8, wantCap: 5 * SizeOfUint8},
		{name: "", args: args{s: make([]uint8, 30, 50)}, want: make([]byte, 30*SizeOfUint8), wantLen: 30 * SizeOfUint8, wantCap: 50 * SizeOfUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTouint8Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint8Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint8Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToUint16Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]uint16, int(1/SizeOfUint16)), wantLen: int(1 / SizeOfUint16), wantCap: int(5 / SizeOfUint16)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]uint16, int(30/SizeOfUint16)), wantLen: int(30 / SizeOfUint16), wantCap: int(50 / SizeOfUint16)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToUint16Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTouint16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint16SliceToBytes(t *testing.T) {
	type args struct {
		s []uint16
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]uint16, 1, 5)}, want: make([]byte, 1*SizeOfUint16), wantLen: 1 * SizeOfUint16, wantCap: 5 * SizeOfUint16},
		{name: "", args: args{s: make([]uint16, 30, 50)}, want: make([]byte, 30*SizeOfUint16), wantLen: 30 * SizeOfUint16, wantCap: 50 * SizeOfUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTouint16Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint16Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint16Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToUint32Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []uint32
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]uint32, int(1/SizeOfUint32)), wantLen: int(1 / SizeOfUint32), wantCap: int(5 / SizeOfUint32)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]uint32, int(30/SizeOfUint32)), wantLen: int(30 / SizeOfUint32), wantCap: int(50 / SizeOfUint32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToUint32Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTouint32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint32SliceToBytes(t *testing.T) {
	type args struct {
		s []uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]uint32, 1, 5)}, want: make([]byte, 1*SizeOfUint32), wantLen: 1 * SizeOfUint32, wantCap: 5 * SizeOfUint32},
		{name: "", args: args{s: make([]uint32, 30, 50)}, want: make([]byte, 30*SizeOfUint32), wantLen: 30 * SizeOfUint32, wantCap: 50 * SizeOfUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTouint32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToUint64Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]uint64, int(1/SizeOfUint64)), wantLen: int(1 / SizeOfUint64), wantCap: int(5 / SizeOfUint64)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]uint64, int(30/SizeOfUint64)), wantLen: int(30 / SizeOfUint64), wantCap: int(50 / SizeOfUint64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToUint64Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTouint64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestUint64SliceToBytes(t *testing.T) {
	type args struct {
		s []uint64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]uint64, 1, 5)}, want: make([]byte, 1*SizeOfUint64), wantLen: 1 * SizeOfUint64, wantCap: 5 * SizeOfUint64},
		{name: "", args: args{s: make([]uint64, 30, 50)}, want: make([]byte, 30*SizeOfUint64), wantLen: 30 * SizeOfUint64, wantCap: 50 * SizeOfUint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTouint64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToUint64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToUint64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToFloat32Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]float32, int(1/SizeOfFloat32)), wantLen: int(1 / SizeOfFloat32), wantCap: int(5 / SizeOfFloat32)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]float32, int(30/SizeOfFloat32)), wantLen: int(30 / SizeOfFloat32), wantCap: int(50 / SizeOfFloat32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToFloat32Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTofloat32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToFloat32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToFloat32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat32SliceToBytes(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]float32, 1, 5)}, want: make([]byte, 1*SizeOfFloat32), wantLen: 1 * SizeOfFloat32, wantCap: 5 * SizeOfFloat32},
		{name: "", args: args{s: make([]float32, 30, 50)}, want: make([]byte, 30*SizeOfFloat32), wantLen: 30 * SizeOfFloat32, wantCap: 50 * SizeOfFloat32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTofloat32Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToFloat32Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToFloat32Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToFloat64Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]float64, int(1/SizeOfFloat64)), wantLen: int(1 / SizeOfFloat64), wantCap: int(5 / SizeOfFloat64)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]float64, int(30/SizeOfFloat64)), wantLen: int(30 / SizeOfFloat64), wantCap: int(50 / SizeOfFloat64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToFloat64Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTofloat64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToFloat64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToFloat64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestFloat64SliceToBytes(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]float64, 1, 5)}, want: make([]byte, 1*SizeOfFloat64), wantLen: 1 * SizeOfFloat64, wantCap: 5 * SizeOfFloat64},
		{name: "", args: args{s: make([]float64, 30, 50)}, want: make([]byte, 30*SizeOfFloat64), wantLen: 30 * SizeOfFloat64, wantCap: 50 * SizeOfFloat64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTofloat64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToFloat64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToFloat64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToComplex64Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []complex64
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]complex64, int(1/SizeOfComplex64)), wantLen: int(1 / SizeOfComplex64), wantCap: int(5 / SizeOfComplex64)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]complex64, int(30/SizeOfComplex64)), wantLen: int(30 / SizeOfComplex64), wantCap: int(50 / SizeOfComplex64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToComplex64Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTocomplex64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToComplex64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToComplex64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex64SliceToBytes(t *testing.T) {
	type args struct {
		s []complex64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]complex64, 1, 5)}, want: make([]byte, 1*SizeOfComplex64), wantLen: 1 * SizeOfComplex64, wantCap: 5 * SizeOfComplex64},
		{name: "", args: args{s: make([]complex64, 30, 50)}, want: make([]byte, 30*SizeOfComplex64), wantLen: 30 * SizeOfComplex64, wantCap: 50 * SizeOfComplex64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Complex64SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTocomplex64Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToComplex64Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToComplex64Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToComplex128Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []complex128
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]complex128, int(1/SizeOfComplex128)), wantLen: int(1 / SizeOfComplex128), wantCap: int(5 / SizeOfComplex128)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]complex128, int(30/SizeOfComplex128)), wantLen: int(30 / SizeOfComplex128), wantCap: int(50 / SizeOfComplex128)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToComplex128Slice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTocomplex128Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToComplex128Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToComplex128Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestComplex128SliceToBytes(t *testing.T) {
	type args struct {
		s []complex128
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]complex128, 1, 5)}, want: make([]byte, 1*SizeOfComplex128), wantLen: 1 * SizeOfComplex128, wantCap: 5 * SizeOfComplex128},
		{name: "", args: args{s: make([]complex128, 30, 50)}, want: make([]byte, 30*SizeOfComplex128), wantLen: 30 * SizeOfComplex128, wantCap: 50 * SizeOfComplex128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Complex128SliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTocomplex128Slice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToComplex128Slice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToComplex128Slice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}
func TestBytesToStringSlice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantLen int
		wantCap int
	}{
		{name: "", args: args{b: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{b: make([]byte, 1, 5)}, want: make([]string, int(1/SizeOfString)), wantLen: int(1 / SizeOfString), wantCap: int(5 / SizeOfString)},
		{name: "", args: args{b: make([]byte, 30, 50)}, want: make([]string, int(30/SizeOfString)), wantLen: int(30 / SizeOfString), wantCap: int(50 / SizeOfString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToStringSlice(tt.args.b)
			l := len(got)
			c := cap(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesTostringSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToStringSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToStringSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

func TestStringSliceToBytes(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantLen int
		wantCap int
	}{
		{name: "", args: args{s: nil}, want: nil, wantLen: 0, wantCap: 0},
		{name: "", args: args{s: make([]string, 1, 5)}, want: make([]byte, 1*SizeOfString), wantLen: 1 * SizeOfString, wantCap: 5 * SizeOfString},
		{name: "", args: args{s: make([]string, 30, 50)}, want: make([]byte, 30*SizeOfString), wantLen: 30 * SizeOfString, wantCap: 50 * SizeOfString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringSliceToBytes(tt.args.s)
			l := len(got)
			c := cap(got)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("BytesTostringSlice() = %v, want %v", got, tt.want)
			}
			if tt.wantLen > 0 {
				if l != tt.wantLen {
					t.Errorf("BytesToStringSlice() len = %v, want %v", l, tt.wantLen)
				}
			}
			if tt.wantCap > 0 {
				if c != tt.wantCap {
					t.Errorf("BytesToStringSlice() cap = %v, want %v", c, tt.wantCap)
				}
			}
		})
	}
}

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
