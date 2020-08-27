package _map

import (
	"testing"
)

func TestGroupInline(t *testing.T) {
	ageNameMap := IterableMap{
		22: "John",
		12: "Bill",
		23: "David",
		14: "Andrew",
	}

	ageNameMap.iterateOver(func(key interface{}, value interface{}, elIdx int) {
		if elIdx == 0 {
			if key != 12 && value.(string) != "Bill" { t.Error() }
		}
		if elIdx == 3 {
			if key != 23 && value.(string) != "David" { t.Error() }
		}
	}, ORDER_ASC)
}


