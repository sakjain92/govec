package blas

import (
	"testing"
)

/* Saxpy */

type saxpyFunc func(N int, alpha float32, X []float32, incX int, Y []float32, incY int)

func helperBenchmarkSaxpy(b *testing.B, fn  saxpyFunc) {

	b.StopTimer()

	var alpha float32

	count := 100000

	f1 := make([]float32, count)
	f2 := make([]float32, count)

	alpha = 2.0

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(count, alpha, f1, 1, f2, 1)
	}
}

func helperTestSaxpy(t *testing.T, fn  saxpyFunc) {

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

	fn(8, alpha, f1, 1, res, 1)

	for i := 0; i < 8; i++ {
		expected := alpha * f1[i] + f2[i]
		if res[i] != expected {
			t.Errorf("Test(): \nexpected %f\ngot      %f", expected, res[i])
		}
	}
}
