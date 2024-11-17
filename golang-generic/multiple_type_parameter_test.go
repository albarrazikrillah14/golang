package golang_generic

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParamter(t *testing.T) {
	MultipleParameter[string, int]("Medomeckz", 1405)
	MultipleParameter[int, string](1405, "Albarra")
}
