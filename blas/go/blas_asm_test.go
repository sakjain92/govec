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

/* Dasum */
func BenchmarkNativeDasum(b *testing.B) {
	helperBenchmarkDasum(b, NativeDasum)
}

func TestNativeDasum(t *testing.T) {
	helperTestDasum(t, NativeDasum)
}

/* Sdot */
func BenchmarkNativeSdot(b *testing.B) {
	helperBenchmarkSdot(b, NativeSdot)
}

func TestNativeSdot(t *testing.T) {
	helperTestSdot(t, NativeSdot)
}

/* Sdsdot */
func BenchmarkNativeSdsdot(b *testing.B) {
	helperBenchmarkSdsdot(b, NativeSdsdot)
}

func TestNativeSdsdot(t *testing.T) {
	helperTestSdsdot(t, NativeSdsdot)
}
