package test

import (
	"fmt"
	"testing"
)

func deferFuncParameter() {
	var aInt = 1
	defer fmt.Println(aInt)
	aInt = 2
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}
func deferFuncParameter2() {
	var aArray = [3]int{1, 2, 3}
	defer printArray(&aArray)
	aArray[0] = 10
}

func deferFuncReturn2() (result int) {
	i := 1
	defer func() {
		result++
	}()
	i = 100
	return i
}

func TestDefer(t *testing.T) {
	deferFuncParameter()

	fmt.Println(deferFuncReturn())

	deferFuncParameter2()
	fmt.Println(deferFuncReturn2())
}
