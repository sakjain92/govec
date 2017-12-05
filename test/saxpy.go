package main

/*
// #include "govec/saxpy.c"
import "C"
*/

/*
// #cgo CFLAGS: -Igovec
// #cgo LDFLAGS: govec/libsaxpy.a
// #include <saxpy.h>
import "C"
*/


// #cgo CFLAGS: -Igovec
// #cgo LDFLAGS: govec/libsaxpy_ispc.a
// #include <saxpy.h>
import "C"


import (
	"fmt"
)

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

/*
func _govecSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {

	var i int

	for i = 0; i < N; i++ {
		Y[i] += alpha * X[i]
	}
}
*/

func saxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) {
	C.govecSaxpy(C.int(N), C.float(alpha), (*C.float)(&X[0]), C.int(incX), (*C.float)(&Y[0]), C.int(incY))
}

func main() {

        var alpha float32
        f1 := []float32{}
        f2 := []float32{}
        res := []float32{}

        alpha = 2.0

        for i := 0; i < 8; i++ {
                f1 = append(f1, float32(i))
                f2 = append(f2, float32(i * 2))
                res = append(res, f2[i])
        }

        saxpy(8, alpha, f1, 1, res, 1)

        for i := 0; i < 8; i++ {
                expected := alpha * f1[i] + f2[i]
                if res[i] != expected {
                        panic("Results don't match")
                }
        }

        fmt.Println("Main done")
}

