//+build !go1.8

package essentials

import "reflect"

func swapper(list interface{}) func(i, j int) {
	val := reflect.ValueOf(list)
	return func(i, j int) {
		val1 := val.Index(i)
		val2 := val.Index(j)
		saved1 := val1.Elem()
		val1.Set(val2.Elem())
		val2.Set(saved1)
	}
}
