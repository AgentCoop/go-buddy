package _map

import (
	"fmt"
	"errors"
	"strings"
	"sort"
)

type IterateOrder uint8

const (
	ORDER_ASC 	IterateOrder = 1
	ORDER_DESC 	IterateOrder = 2
)

const (
	err_WrongIterationOrder = "wrong iteration order"
)

type IterableMap map[interface{}]interface{}

type MapIterator func(interface{}, interface{}, int)

func (m IterableMap) IterateOver(f MapIterator, order IterateOrder) {
	var intKeys []int
	var strKeys []string
	for k := range m {
		switch T := k.(type) {
		case int:
			intKeys = append(intKeys, k.(int))
		case string:
			strKeys = append(strKeys, k.(string))
		default:
			panic(errors.New(fmt.Sprintf("iterable map: unsupported key type %v", T)))
		}
	}
	if len(intKeys) > 0 {
		sort.Slice(intKeys, func(i int, j int) bool {
			a, b := intKeys[i], intKeys[j]
			if order == ORDER_ASC {
				return a < b
			} else if order == ORDER_DESC {
				return a > b
			} else {
				panic(err_WrongIterationOrder)
			}
		})
		elIdx := 0
		for _, key := range intKeys {
			f(key, m[key], elIdx)
			elIdx++
		}
	} else if len(strKeys) > 0 {
		sort.Slice(strKeys, func(i int, j int) bool {
			cmp := strings.Compare(strKeys[i], strKeys[j])
			if cmp == 0 {
				return false
			} else if order == ORDER_ASC {
				return cmp == -1
			} else if order == ORDER_DESC {
				return cmp != -1
			} else {
				panic(err_WrongIterationOrder)
			}
		})
		elIdx := 0
		for _, key := range strKeys {
			f(key, m[key], elIdx)
			elIdx++
		}
	}
}
