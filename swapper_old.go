//+build !go1.8

package essentials

import "reflect"

func swapper(list interface{}) func(i, j int) {
	val := reflect.ValueOf(list)
	return func(i, j int) {
		val1 := val.Index(i)
		val2 := val.Index(j)
		backup := val1.Interface()
		val1.Set(val2)
		val2.Set(reflect.ValueOf(backup))
	}
}
