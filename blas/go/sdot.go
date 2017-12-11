//+build !ASSEMBLY

package blas

import (
	"github.com/sakjain92/govectool/govec"
)

//go:generate govectool sdot.go

func SerialSdot(N int, X []float32, incX int, Y []float32, incY int) float32 {
	var (
		sum float32
		i int
	)

	if incX != 1 && incY != 1 {
		panic("Wrong arguments")
	}

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

func _govecISPCSdot(N govec.UniformInt, X []govec.UniformFloat32,
			incX govec.UniformInt,
			Y []govec.UniformFloat32, incY govec.UniformInt) govec.UniformFloat32 {
	var sum float32

	for i := range govec.Range(0, N) {
		sum += (float32)(X[i] * Y[i])
	}
	return  govec.UniformFloat32(govec.ReduceAddFloat32(sum))
}
