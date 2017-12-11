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
	for ; N > 0; N-- {
		a += float64(X[xi]) * float64(Y[yi])
		xi += incX
		yi += incY
	}
	return float32(float64(alpha) + a)
}

func SerialSdsdot(N int, alpha float32, X []float32, Y []float32) float32 {
	var a float64
	var i int
	for ; i < N; i++ {
		a += float64(X[i]) * float64(Y[i])
	}
	return float32(float64(alpha) + a)
}

func _govecSdsdot(N govec.UniformInt, alpha govec.UniformFloat32,
				  X []govec.UniformFloat32, Y []govec.UniformFloat32) govec.UniformFloat32 {
	var sum float64
	for i := range govec.Range(0, N) {
		sum += (float64)(X[i]) * (float64)(Y[i])
	}
	return (govec.UniformFloat32)((float64)(alpha) + govec.ReduceAddFloat64(sum))
}
