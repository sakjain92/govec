package saxpy

/*
// #include "govec/saxpy.c"
import "C"
*/


// #cgo CFLAGS: -Igovec_build
// #cgo LDFLAGS: govec_build/libsaxpy.a
// #include <saxpy.h>
import "C"


/*
// #cgo CFLAGS: -Igovec
// #cgo LDFLAGS: govec/libsaxpy_ispc.a
// #include <saxpy.h>
import "C"
*/

import (
	"github.com/sakjain92/govectool/govec"
)



func GoSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {

	for i := 0; i < N; i++ {
		Y[i] += alpha * X[i]
	}
}


func _govecSaxpy(N govec.UniformInt, alpha govec.UniformFloat32,
		X []govec.UniformFloat32, incX govec.UniformInt,
		 Y []govec.UniformFloat32, incY govec.UniformInt) {

	for i := range govec.Range(0, N) {
		Y[i] += alpha * X[i]
	}
}


/*
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
*/

/*
func _govecSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {

	var i int

	for i = 0; i < N; i++ {
		Y[i] += alpha * X[i]
	}
}
*/

func CSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {
	C.govecSaxpy(C.int(N), C.float(alpha), (*C.float)(&X[0]), C.int(incX), (*C.float)(&Y[0]), C.int(incY))
}
