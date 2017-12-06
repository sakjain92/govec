package saxpy

import (
	"testing"
	"fmt"
)

func TestSaxpy(t *testing.T) {

	var alpha float32
	f1 := []float32{}
	f2 := []float32{}
	res := []float32{}

	alpha = 2.0

	for i := 0; i < 8; i++ {
		f1 = append(f1, float32(i))
		f2 = append(f2, float32(i * 2))
		res = append(res, f2[i])
	}

	Saxpy(8, alpha, f1, 1, res, 1)

	for i := 0; i < 8; i++ {
		expected := alpha * f1[i] + f2[i]
		if res[i] != expected {
			t.Errorf("Test(): \nexpected %f\ngot      %f", expected, res[i])
		}
	}

	fmt.Println("TestSaxpy Passed")
}
