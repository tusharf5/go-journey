package hmap

import (
	"fmt"
)

type HMap struct {
	arr  []interface{}
	cap  int
	keys []string
}

func New() *HMap {
	arr := make([]interface{}, 20)
	return &HMap{arr, 20, make([]string, 0)}
}

func (d *HMap) Set(key string, value interface{}) {

	d.keys = append(d.keys, key)
	byteSum := 0

	for _, val := range key {
		byteSum = byteSum + int(val)
	}

	index := byteSum % d.cap

	d.arr[index] = value

	if len(d.keys) == d.cap {
		// if array size is full re-hash
		fmt.Println("Re-Hashing")
		newCap := cap(d.arr) + 20
		newArr := make([]interface{}, newCap)
		d.cap = newCap

		for _, keyVal := range d.keys {
			fmt.Println("Re-Hashing", string(keyVal))

			internalbyteSum := 0

			for _, val := range keyVal {
				internalbyteSum = internalbyteSum + int(val)
			}

			newIndex := internalbyteSum % newCap

			newArr[newIndex] = value
		}

		d.arr = newArr
	}

}

func (d *HMap) Get(key string) interface{} {
	byteSum := 0

	for _, val := range key {
		byteSum = byteSum + int(val)
	}

	index := byteSum % d.cap

	return d.arr[index]

}
