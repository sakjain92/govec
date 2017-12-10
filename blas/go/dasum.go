package dasum

import (
	"github.com/sakjain92/govectool/govec"
)

func NativeDasum(N int, X []float64, incX int) float64

func SerialDasum(N int, X []float64) float64 {
	var (
		a  float64
		i int
	)
	for ; N > 0; N-- {
		x := X[i]
		if x < 0 {
			x = -x
		}
		a += x
		i++
	}
	return a
}

func SerialDasumGeneric(N int, X []float64, incX int) float64 {
	var (
		a  float64
		xi int
	)
	for ; N > 0; N-- {
		x := X[xi]
		if x < 0 {
			x = -x
		}
		a += x
		xi += incX
	}
	return a
}

func _govecDasum(N govec.UniformInt, X []govec.UniformFloat64, 
				 incX govec.UniformInt) govec.UniformFloat64 {
	var sum, x float64

	for i := range govec.Range(0, N) {
		x = X[i]
		if x < 0 {
			x = -x
		}
		sum += x
	}
	return govec.ReduceAdd(sum)
}