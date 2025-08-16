package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Inside Main")

	// fmt.Println(deepcopy.DeepCopy("Hello"))

	// x := 4
	// y := &x
	// fmt.Println(reflect.ValueOf(&x).Elem().Addr())
	// fmt.Println(reflect.ValueOf(y).Elem().Addr())

	// var y interface{} = nil
	// fmt.Println(reflect.ValueOf(y)) // <nil>
	// fmt.Println(reflect.TypeOf(y)) // <nil>
	// fmt.Println(reflect.ValueOf(y).Type())
	// fmt.Println(reflect.ValueOf(y).Kind())


	var x *int
	v:=reflect.ValueOf(x)
	t:=reflect.TypeOf(x)
	fmt.Println(v)
	fmt.Println(v.Elem())
	fmt.Println(t.Elem())
	fmt.Println(v.IsNil())
	dc := reflect.New(t.Elem())
	fmt.Println(dc==reflect.Zero(t))
	fmt.Println(dc)
	fmt.Println(reflect.Zero(t))
	// fmt.Println(v.Elem().Addr())
	// fmt.Println(v.Elem().Addr().Interface())
	// fmt.Println(v.Elem().Addr().Interface().(*int))
	// fmt.Println(v.Elem().Addr().Interface().(*int) == nil)

}
