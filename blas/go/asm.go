//+build ASSEMBLY

package blas


func NativeSaxpy(N int, alpha float32, X []float32, incX int, Y []float32,
				 incY int)

func NativeSdot(N int, X []float32, incX int, Y []float32, incY int) float32

func NativeSdsdot(N int, alpha float32, X []float32, incX int, Y []float32, incY int) float32

func NativeDasum(N int, X []float64, incX int) float64

