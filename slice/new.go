package slice

import (
	"reflect"
	"unsafe"
)

func (s Sequence) NewInt(seq []int) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Int}
}

func (s Sequence) NewInt8(seq []int8) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Int8}
}

func (s Sequence) NewInt16(seq []int16) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Int16}
}

func (s Sequence) NewInt32(seq []int32) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Int32}
}

func (s Sequence) NewInt64(seq []int64) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Int64}
}

func (s Sequence) NewUint(seq []uint) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Uint}
}

func (s Sequence) NewUint8(seq []uint8) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Uint8}
}

func (s Sequence) NewUint16(seq []uint16) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Uint16}
}

func (s Sequence) NewUint32(seq []uint32) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Uint32}
}

func (s Sequence) NewUint64(seq []uint64) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Uint64}
}

func (s Sequence) NewFloat32(seq []float32) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Float32}
}

func (s Sequence) NewFloat64(seq []float64) Sequence {
	return Sequence{ptr: unsafe.Pointer(&seq), kind: reflect.Float64}
}
