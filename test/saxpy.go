package main

// #include "govec/saxpy.c"
import "C"

func _govecSaxpy(N uniform_int, alpha uniform_float32, X []uniform_float32, Y []uniform_float32) {

	var i uniform_int

	for i = range govecRange(0, N) {
		Y[i] += alpha * X[i]
	}
}

func _govecSaxpy(N uniform_int, alpha uniform_float32, X []uniform_float32,
				incX uniform_int, Y []uniform_float32, incY uniform_int) {

	var xi, yi, i uniform_int

	for i = 0; i < N; i += programCount {
		var index, xi_index, yi_index int

		index = i + programIndex
		xi_index = xi + programIndex * incX 
		yi_index = yi + programIndex * incY 

		Y[yi_index] += alpha * X[xi_index]
		xi += programCount * incX
		yi += programCount *incY
	}
}

func saxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {
	C.govecSaxpy(C.int(N), C.float(alpha), (*C.float)(&X[0]), C.int(incX), (*C.float)(&Y[0]), C.int(incY))
}
