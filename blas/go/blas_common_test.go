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
	exp := []float32{}

	alpha = 2.0

	for i := 0; i < 8; i++ {
		f1 = append(f1, float32(i))
		f2 = append(f2, float32(i * 2))
		exp = append(exp, (alpha * float32(i) )+ (float32(i) * 2))
		res = append(res, f2[i])
	}

	fn(8, alpha, f1, 1, res, 1)

	for i := 0; i < 8; i++ {
		if res[i] != exp[i] {
			t.Errorf("Test(): \nexpected %f\ngot      %f", exp[i], res[i])
		}
		t.Logf("%d:%f\n", i, res[i])
	}
}

/* Mandlebrot */

type mandlebrotFunc func(x0 float32, y0 float32, x1 float32, y1 float32,
			width int, height int, startRow int, totalRows int,
			maxIterations int, output []int32)

func helperBenchmarkMandelbrot(t *testing.B, fn mandlebrotFunc) {

	var width int = 1200;
	var height int = 800;
	var maxIterations int = 256;

	var x0 float32 = -2;
	var x1 float32 = 1;
	var y0 float32 = -1;
	var y1 float32 = 1;

	out := make([]int32, width * height)

	fn(x0, y0, x1, y1, width, height, 0, height, maxIterations,
		out)
}
// Comparing two different implementations of mandlebrot
func helperTestMandelbrot(t *testing.T, fn1 mandlebrotFunc,  fn2 mandlebrotFunc) {

	var width int = 1200;
	var height int = 800;
	var maxIterations int = 256;

	var x0 float32 = -2;
	var x1 float32 = 1;
	var y0 float32 = -1;
	var y1 float32 = 1;

	out1 := make([]int32, width * height)
	out2 := make([]int32, width * height)

	fn1(x0, y0, x1, y1, width, height, 0, height, maxIterations,
		out1)

	fn2(x0, y0, x1, y1, width, height, 0, height, maxIterations,
		out2)

	if len(out1) != len(out2) {
		t.Errorf("Length don't match: %d, %d", len(out1), len(out2))
	}

	for i := range(out1) {
		if out1[i] != out2[i] {
			t.Errorf("Results don't match for index %d: %d, %d",
				i, out1[i], out2[i])
		}
	}
}
