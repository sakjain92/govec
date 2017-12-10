package saxpy

import (
	"github.com/sakjain92/govectool/govec"
)

func SerialSaxpy(N int, alpha float32, X []float32, Y []float32) {

	for i := 0; i < N; i++ {
		Y[i] += alpha * X[i]
	}
}

func SerialSaxpyGeneric(N int, alpha float32, X []float32, incX int,
						Y []float32, incY int) {

	if incX == 1 && incY == 1 {
		for i := 0; i < N; i++ {
			Y[i] += alpha * X[i]
		}
		return
	}

	var xi, yi int
	switch alpha {
	case 0:
	case 1:
		for ; N >= 2; N -= 2 {
			Y[yi] += X[xi]
			xi += incX
			yi += incY

			Y[yi] += X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	case -1:
		for ; N >= 2; N -= 2 {
			Y[yi] -= X[xi]
			xi += incX
			yi += incY

			Y[yi] -= X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] -= X[xi]
		}
	default:
		for ; N >= 2; N -= 2 {
			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY

			Y[yi] += alpha * X[xi]
			xi += incX
			yi += incY
		}
		if N != 0 {
			Y[yi] += alpha * X[xi]
		}
	}
}

func NativeSaxpy(N int, alpha float32, X []float32, incX int, Y []float32,
				 incY int)

func _govecSaxpy(N govec.UniformInt, alpha govec.UniformFloat32,
				 X []govec.UniformFloat32, Y []govec.UniformFloat32) {

	for i := range govec.Range(0, N) {
		Y[i] += alpha * X[i]
	}
}

func _govecSaxpyGeneric(N govec.UniformInt, alpha govec.UniformFloat32,
				 X []govec.UniformFloat32, incX govec.UniformInt,
				 Y []govec.UniformFloat32, incY govec.UniformInt) {

	var xi, yi, i govec.UniformInt

	for i = 0; i < N; i += govec.ProgramCount {
		var index, xi_index, yi_index int

		index = i + govec.ProgramIndex
		xi_index = xi + govec.ProgramIndex * incX
		yi_index = yi + govec.ProgramIndex * incY

		Y[yi_index] += alpha * X[xi_index]
		xi += govec.ProgramCount * incX
		yi += govec.ProgramCount *incY
	}
}
