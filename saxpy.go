package main

func govecSaxpy(N int, alpha float32, X []float32, incX int, Y []float32, incY int) (int) {

	var xi, yi, i int

	for i = 0; i < N; i++ {
		Y[yi] += alpha * X[xi]
		xi += incX
		yi += incY
	}

	return 0
}
