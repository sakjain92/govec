package sdot

import (
	"github.com/sakjain92/govectool/govec"
)

// Scalar product: X^T Y
func NativeSdot(N int, X []float32, incX int, Y []float32, incY int) float32

func SerialSdot(N int, X []float32, Y []float32) float32 {
	var (
		sum float32
		i int
	)
	for ; i < N; i++  {
		sum += X[i] * Y[i]
	}
	return sum
}

func SerialSdotGeneric(N int, X []float32, incX int, Y []float32, incY int) float32 {
	var (
		sum float32
		xi, yi     int
	)
	for ; N > 0; N-- {
		sum += X[xi] * Y[yi]
		xi += incX
		yi += incY
	}
	return sum
}

func _govecSdot(N govec.UniformInt, X []govec.UniformFloat32,
				Y []govec.UniformFloat32) govec.UniformFloat32 {
	var sum float32

	for i := range govec.Range(0, N) {
		sum += X[i] * Y[i]
	}
	return govec.ReduceAdd(sum)
}
