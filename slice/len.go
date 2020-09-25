package slice

func (s Sequence) Len() int {
	slice := *(*[]interface{})(s.ptr)
	return len(slice)
}