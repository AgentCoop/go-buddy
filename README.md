# go-buddy
Helpers methods and functions for builtin Go types

## Map

Iterate over a map in the specified order
```go
	ageNameMap := IterableMap{
		22: "John",
		12: "Bill",
		23: "David",
		14: "Andrew",
	}
	ageNameMap.iterateOver(func(key interface{}, val interface{}, elIdx int) {
		fmt.Printf("Age: %d, Name: %s index: %d\n", key.(int), val, elIdx)
	}, ORDER_ASC)
```