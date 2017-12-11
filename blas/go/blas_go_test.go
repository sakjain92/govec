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

/* Dasum */

func BenchmarkDasumSaxpy(b *testing.B) {
	helperBenchmarkDasum(b, SerialDasum)
}

func BenchmarkSerialDasumGeneric(b *testing.B) {
	helperBenchmarkDasum(b, SerialDasumGeneric)
}

func BenchmarkISPCDasum(b *testing.B) {
	helperBenchmarkDasum(b, ISPCDasum)
}


func TestSerialDasum(t *testing.T) {
	helperTestDasum(t, SerialDasum)
}

func TestSerialDasumGeneric(t *testing.T) {
	helperTestDasum(t, SerialDasumGeneric)
}

func TestISPCDasum(t *testing.T) {
	helperTestDasum(t, ISPCDasum)
}

/* Sdot */

func BenchmarkSdotSaxpy(b *testing.B) {
	helperBenchmarkSdot(b, SerialSdot)
}

func BenchmarkSerialSdotGeneric(b *testing.B) {
	helperBenchmarkSdot(b, SerialSdotGeneric)
}

func BenchmarkISPCSdot(b *testing.B) {
	helperBenchmarkSdot(b, ISPCSdot)
}

func TestSerialSdot(t *testing.T) {
	helperTestSdot(t, SerialSdot)
}

func TestSerialSdotGeneric(t *testing.T) {
	helperTestSdot(t, SerialSdotGeneric)
}

func TestISPCSdot(t *testing.T) {
	helperTestSdot(t, ISPCSdot)
}

/* Sdsdot */

func BenchmarkSdsdotSaxpy(b *testing.B) {
	helperBenchmarkSdsdot(b, SerialSdsdot)
}

func BenchmarkSerialSdsdotGeneric(b *testing.B) {
	helperBenchmarkSdsdot(b, SerialSdsdotGeneric)
}

func BenchmarkISPCSdsdot(b *testing.B) {
	helperBenchmarkSdsdot(b, ISPCSdsdot)
}

func TestSerialSdsdot(t *testing.T) {
	helperTestSdsdot(t, SerialSdsdot)
}

func TestSerialSdsdotGeneric(t *testing.T) {
	helperTestSdsdot(t, SerialSdsdotGeneric)
}

func TestISPCSdsdot(t *testing.T) {
	helperTestSdsdot(t, ISPCSdsdot)
}

/* Mandlebrot */
func BenchmarkSerialMandelBrot(b *testing.B) {
	helperBenchmarkMandelbrot(b, SerialMandelbrot)
}

func BenchmarkISPCMandelBrot(b *testing.B) {
	helperBenchmarkMandelbrot(b, ISPCMandelbrot)
}

func TestBothMandleBrot(t *testing.T) {
	helperTestMandelbrot(t, SerialMandelbrot, ISPCMandelbrot)
}

/* Sleep */
func BenchmarkSleep(b *testing.B) {

	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second)
	}
}


