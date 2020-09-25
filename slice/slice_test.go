package slice_test

import (
	. "github.com/AgentCoop/go-buddy/slice"
	"testing"
)

func getFirstItem(s Slice) interface{} {
	return s.First()
}

func getLastItem(s Slice) interface{} {
	return s.Last()
}

func TestSlice(t *testing.T) {
	var nums = Sequence{}.NewInt16([]int16{0, 1, 3, 5, 8})
	var primes = nums.Slice(1, 3)
	if primes.Len() != 3 { t.Fatalf(".Slice expected 3, got: %d\n", primes.Len()) }
	if getFirstItem(primes).(int16) != 1 { t.Fatalf(".First expected 1, got: %d\n", primes.First()) }
	if getLastItem(primes).(int16) != 5 { t.Fatalf(".Last expected 5, got: %d\n", primes.Last()) }

	var one = primes.Slice(1, 1)
	if one.Len() != 1 {  t.Fatalf(".Slice(1,1) %d\n", primes.First()) }

	var negLen = nums.Slice(1, -2)
	if negLen.Len() != 2 {  t.Fatalf(".Slice(1,-2)") }

	negLen = nums.Slice(1, -9)
	if negLen.Len() != 0 {  t.Fatalf(".Slice(1,-9)") }
}
