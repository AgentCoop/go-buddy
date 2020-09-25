package slice

import "reflect"

func (s Sequence) First() interface{} {
	if s.Len() < 1 { return nil }
	m := typeHandlerMapI{
		reflect.Int: func() interface{} {
			slice := *(*[]int)(s.ptr)
			return slice[0]
		},
		reflect.Int16: func() interface{} {
			slice := *(*[]int16)(s.ptr)
			return slice[0]
		},
	}
	return callHandlerI(s.kind, m)
}
