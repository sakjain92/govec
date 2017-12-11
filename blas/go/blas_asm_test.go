//+build ASSEMBLY

package blas

import (
	"testing"
)

/* Saxpy */
func BenchmarkNativeSaxpy(b *testing.B) {
	helperBenchmarkSaxpy(b, NativeSaxpy)
}

func TestNativeSaxpy(t *testing.T) {
	helperTestSaxpy(t, NativeSaxpy)
}
