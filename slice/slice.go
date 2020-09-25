package slice

import (
	"reflect"
	"runtime"
	"unsafe"
)

// copyWithLeftPad copies src to the end of dest, padding with zero bytes as
// needed.
//func copyWithLeftPad(dest, src []byte) {
//	numPaddingBytes := len(dest) - len(src)
//	for i := 0; i < numPaddingBytes; i++ {
//		dest[i] = 0
//	}
//	copy(dest[numPaddingBytes:], src)
//}

type Sequence struct {
	ptr unsafe.Pointer
	kind reflect.Kind
}

type typeHandlerFunc func() Sequence
type typeHandlerFuncI func() interface{}

type typeHandlerMap map[reflect.Kind]typeHandlerFunc
type typeHandlerMapI map[reflect.Kind]typeHandlerFuncI

type Slice interface {
	NewInt(seq []int) Sequence
	NewInt8(seq []int8) Sequence
	NewInt16(seq []int16) Sequence
	NewInt32(seq []int32) Sequence
	NewInt64(seq []int64) Sequence
	NewUint(seq []uint) Sequence
	NewUint8(seq []uint8) Sequence
	NewUint16(seq []uint16) Sequence
	NewUint32(seq []uint32) Sequence
	NewUint64(seq []uint64) Sequence
	NewFloat32(seq []float32) Sequence
	NewFloat64(seq []float64) Sequence
	Slice(offset, length int) Sequence
	First() interface{}
	Last() interface{}
	Len() int
}

func chkTypeSupport(ok bool) {
	if ok != true {
		_, fn, line, _ := runtime.Caller(1)
		panic("Slice: unsupported type. File: " + fn + "; Line: " + string(line))
	}
}

func callHandler(typ reflect.Kind, hm typeHandlerMap) Sequence {
	h, ok := hm[typ]
	chkTypeSupport(ok)
	return h()
}

func callHandlerI(typ reflect.Kind, hm typeHandlerMapI) interface{} {
	h, ok := hm[typ]
	chkTypeSupport(ok)
	return h()
}

//

func (s Sequence) Slice(offset, length int) Sequence {
	var start, end int
	if offset < 0 { start = -1 * offset } else { start = offset }
	if length < 0 {
		end = s.Len() - -1 * length
		if end <= 0 {
			start = 0
			end = 0
		}
	} else {
		switch {
		case start + length < s.Len():
			end = start + length
		default:
			end = s.Len()
		}
	}
	m := typeHandlerMap{
		reflect.Int: func() Sequence {
			slice := *(*[]int)(s.ptr)
			return s.NewInt(slice[start:end])
		},
		reflect.Int16: func() Sequence {
			slice := *(*[]int16)(s.ptr)
			return s.NewInt16(slice[start:end])
		},
	}
	return callHandler(s.kind, m)
}

func (s Sequence) Val() {
	reflect.ValueOf(s.kind)
}

