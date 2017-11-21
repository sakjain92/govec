package main

// #include "govec/saxpy.c"
import "C"

func _govecSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {

	var xi, yi, i int

	for i = 0; i < N; i++ {
		Y[yi] += alpha * X[xi]
		xi += incX
		yi += incY
	}
}

func saxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {
	C.govecSaxpy(C.int(N), C.float(alpha), (*C.float)(&X[0]), C.int(incX), (*C.float)(&Y[0]), C.int(incY))
}
