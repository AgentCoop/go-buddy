package _map

import (
	"testing"
)

func TestIterator(t *testing.T) {
	ageNameMap := IterableMap{
		22: "John",
		12: "Bill",
		23: "David",
		14: "Andrew",
	}

	nameAgeMap := IterableMap{
		"John": 22,
		"Bill": 12,
		"David": 23,
		"Andrew": 14,
	}

	ageNameMap.IterateOver(func(key interface{}, value interface{}, elIdx int) {
		if elIdx == 0 { if key != 12 && value.(string) != "Bill" { t.Error() } }
		if elIdx == 3 { if key != 23 && value.(string) != "David" { t.Error() } }
	}, ORDER_ASC)

	nameAgeMap.IterateOver(func(key interface{}, value interface{}, elIdx int) {
		if elIdx == 0 { if key.(string) != "John" { t.Error() } }
		if elIdx == 3 {if key.(string) != "Andrew" { t.Error() } }
	}, ORDER_DESC)
}


