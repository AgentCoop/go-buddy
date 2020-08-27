# go-buddy
Helper methods and functions for builtin Go types.

## Usage

#### Map
Iterate over map keys in the specified order:
```go
import (
    _m "github.com/AgentCoop/go-buddy/_map"
)

ageNameMap := _m.IterableMap{
    22: "John",
    12: "Bill",
    23: "David",
    14: "Andrew",
}
ageNameMap.iterateOver(func(key interface{}, val interface{}, elIdx int) {
    fmt.Printf("Age: %d, Name: %s, index: %d\n", key.(int), val, elIdx)
}, _m.ORDER_ASC)
```