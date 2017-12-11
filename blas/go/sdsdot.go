//+build !ASSEMBLY

package blas

//go:generate govectool sdsdot.go

import (
	"github.com/sakjain92/govectool/govec"
)

func SerialSdsdotGeneric(N int, alpha float32, X []float32, incX int, Y []float32, incY int) float32 {
	var (
		a float64
		xi, yi     int
	)

	if incX != 1 && incY != 1 {
		panic("Wrong arguments")
	}

	for ; N > 0; N-- {
		a += float64(X[xi]) * float64(Y[yi])
		xi += incX
		yi += incY
	}
	return float32(float64(alpha) + a)
}

func SerialSdsdot(N int, alpha float32, X []float32, incX int, Y []float32, incY int) float32 {
	var a float64
	var i int
	for ; i < N; i++ {
		a += float64(X[i]) * float64(Y[i])
	}
	return float32(float64(alpha) + a)
}

func _govecISPCSdsdot(N govec.UniformInt, alpha govec.UniformFloat32,
			X []govec.UniformFloat32, incX govec.UniformInt,
			Y []govec.UniformFloat32, incY govec.UniformInt) govec.UniformFloat32 {
	var sum float64
	for i := range govec.Range(0, N) {
		sum += (float64)(X[i]) * (float64)(Y[i])
	}
	return (govec.UniformFloat32)((float64)(alpha) + govec.ReduceAddFloat64(sum))
}
