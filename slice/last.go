package slice

import "reflect"

func (s Sequence) Last() interface{} {
	if s.Len() < 1 { return nil }
	m := typeHandlerMapI{
		reflect.Int: func() interface{} {
			slice := *(*[]int)(s.ptr)
			return slice[len(slice) - 1]
		},
		reflect.Int16: func() interface{} {
			slice := *(*[]int16)(s.ptr)
			return slice[len(slice) - 1]
		},
	}
	return callHandlerI(s.kind, m)
}

//func (s SliceInt16) Last() int16 {
//	return s[s.Len() - 1]
//}