package deepcopy

import (
	"fmt"
	"testing"
	"reflect"
	"strconv"
)

func TestNestedStructs(t *testing.T) {
	type Inner struct {
		C string
		D []int
		E map[string]int
		F *Inner
	}
	type Outer struct {
		A int
		B *Outer
		I1 *Inner
		I2 *Inner
	}

	Obj1 := Inner{
		C: "Hello",
		D: []int{1, 2, 3},
		E: map[string]int{"a": 1, "b": 2, "c": 3},
	}
	Obj1.F = &Obj1

	Obj3 := Outer{
		A: 1,
	}
	Obj3.B = &Obj3
	Obj3.I1 = &Obj1
	Obj3.I2 = &Obj1

	tests := []any{Obj3}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			new, err := DeepCopy(test)
			if err != nil {
				t.Errorf("deep copy failed with error %v", err)
			}

			fmt.Println(&test, &new)
			fmt.Println(reflect.DeepEqual(test, new))
			fmt.Println(reflect.DeepEqual(test, test))
			fmt.Println(test,test.(Outer).B,test.(Outer).I1,test.(Outer).I2)
			fmt.Println(new,new.(Outer).B,new.(Outer).I1,new.(Outer).I2)	

			if test.(Outer).B == new.(Outer).B {
				t.Errorf("deep copy failed both old and new have same address")
			}
			if test.(Outer).I1 == new.(Outer).I1 {
				t.Errorf("deep copy failed both old and new have same address")
			}
			if test.(Outer).I2 == new.(Outer).I2 {
				t.Errorf("deep copy failed both old and new have same address")
			}
			if test.(Outer).I1.F == new.(Outer).I1.F {
				t.Errorf("deep copy failed both old and new have same address")
			}
			if !reflect.DeepEqual(test, new) {
				t.Errorf("deep copy failed")
			}
		})
	}

}
