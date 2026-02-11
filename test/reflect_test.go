package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var x float64 = 3.4
	t1 := reflect.TypeOf(x) //t is reflext.Type
	fmt.Println("type:", t1)
	v := reflect.ValueOf(x) //v is reflext.Value
	fmt.Println("value:", v)
}
