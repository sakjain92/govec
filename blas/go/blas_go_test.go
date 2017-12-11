//+build !ASSEMBLY

package blas

import (
	"testing"
	"time"
)

/* Saxpy */

func BenchmarkSerialSaxpy(b *testing.B) {
	helperBenchmarkSaxpy(b, SerialSaxpy)
}

func BenchmarkSerialSaxpyGeneric(b *testing.B) {
	helperBenchmarkSaxpy(b, SerialSaxpyGeneric)
}

func BenchmarkISPCSaxpy(b *testing.B) {
	helperBenchmarkSaxpy(b, ISPCSaxpy)
}

func BenchmarkISPCSaxpyGeneric(b *testing.B) {
	helperBenchmarkSaxpy(b, ISPCSaxpyGeneric)
}

func BenchmarkSleep(b *testing.B) {
	time.Sleep(1 * time.Second)
}

func TestSerialSaxpy(t *testing.T) {
	helperTestSaxpy(t, SerialSaxpy)
}

func TestSerialSaxpyGeneric(t *testing.T) {
	helperTestSaxpy(t, SerialSaxpyGeneric)
}

func TestISPCSaxpy(t *testing.T) {
	helperTestSaxpy(t, ISPCSaxpy)
}

func TestISPCSaxpyGeneric(t *testing.T) {
	helperTestSaxpy(t, ISPCSaxpyGeneric)
}
