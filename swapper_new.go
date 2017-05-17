//+build go1.8

package essentials

import "reflect"

func swapper(list interface{}) func(i, j int) {
	return reflect.Swapper(list)
}
