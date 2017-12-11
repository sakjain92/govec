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

/* Dasum */

type dasumFunc func(N int, X []float64, incX int) float64

func helperBenchmarkDasum(b *testing.B, fn dasumFunc) {

	b.StopTimer()

	count := 100000

	X := make([]float64, count)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(count, X, 1)
	}
}

func helperTestDasum(t *testing.T, fn dasumFunc) {

	X := []float64{}
	var res float64

	for i := 0; i < 8; i++ {
		if i % 2 == 0 {
			X = append(X, float64(i))
		} else {
			X = append(X, float64(i) * -1)
		}

		res = res + float64(i)
	}

	ans := fn(8, X, 1)

	if res != ans {
		t.Errorf("Test(): \nexpected %f got %f\n", ans, res)
	}
}

/* Sdot */

type sdotFunc func(N int, X []float32, incX int, Y []float32, incY int) float32

func helperBenchmarkSdot(b *testing.B, fn sdotFunc) {

	b.StopTimer()

	count := 100000

	X := make([]float32, count)
	Y := make([]float32, count)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(count, X, 1, Y, 1)
	}
}

func helperTestSdot(t *testing.T, fn sdotFunc) {

	X := []float32{}
	Y := []float32{}

	var res float32

	for i := 0; i < 8; i++ {
		X = append(X, float32(i))
		Y = append(Y, float32(i * 2))
		res = res + float32(i * i * 2)
	}

	ans := fn(8, X, 1, Y, 1)

	if res != ans {
		t.Errorf("Test(): \nexpected %f got %f\n", ans, res)
	}
}

/* Sdsdot */

type sdsdotFunc func(N int, alpha float32, X []float32, incX int, Y []float32, incY int) float32

func helperBenchmarkSdsdot(b *testing.B, fn sdsdotFunc) {
	b.StopTimer()

	count := 100000

	X := make([]float32, count)
	Y := make([]float32, count)
	var alpha float32 = 2.0

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(count, alpha, X, 1, Y, 1)
	}
}

func helperTestSdsdot(t *testing.T, fn sdsdotFunc) {

	X := []float32{}
	Y := []float32{}
	var alpha float32 = 2.0
	var res float32

	for i := 0; i < 8; i++ {
		X = append(X, float32(i))
		Y = append(Y, float32(i * 2))
		res = res + float32(i * i * 2)
	}

	res = res + alpha

	ans := fn(8, alpha, X, 1, Y, 1)

	if res != ans {
		t.Errorf("Test(): \nexpected %f got %f\n", ans, res)
	}

}

/* Mandlebrot */

type mandlebrotFunc func(x0 float32, y0 float32, x1 float32, y1 float32,
			width int, height int, startRow int, totalRows int,
			maxIterations int, output []int32)

func helperBenchmarkMandelbrot(b *testing.B, fn mandlebrotFunc) {

	b.StopTimer()

	var width int = 120;
	var height int = 80;
	var maxIterations int = 256;

	var x0 float32 = -2;
	var x1 float32 = 1;
	var y0 float32 = -1;
	var y1 float32 = 1;

	out := make([]int32, width * height)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(x0, y0, x1, y1, width, height, 0, height, maxIterations,
		out)
	}
}
// Comparing two different implementations of mandlebrot
func helperTestMandelbrot(t *testing.T, fn1 mandlebrotFunc,  fn2 mandlebrotFunc) {

	var width int = 120;
	var height int = 80;
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
