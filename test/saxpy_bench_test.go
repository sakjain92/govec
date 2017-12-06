package saxpy

import "testing"

const (
	count = 10000
)

var result []float32

func BenchmarkCSaxpy(b *testing.B) {

	var alpha float32
	f1 := make([]float32, count)
	f2 := make([]float32, count)

	alpha = 2.0

	for i := 0; i < b.N; i++ {
		Saxpy(count, alpha, f1, 1, f2, 1)
	}

	result = f2

}

func BenchmarkGoSaxpy(b *testing.B) {

	var alpha float32
	f1 := make([]float32, count)
	f2 := make([]float32, count)

	alpha = 2.0


	for i := 0; i < b.N; i++ {
		GoSaxpy(count, alpha, f1, 1, f2, 1)
	}
}
