#include "benchmark.h"

float SerialSdot(int N, float X[], int incX, float Y[], int incY) {
	float sum = 0.0;

	if incX != 1 && incY != 1 {
		return -1;
	}

	for (int i = 0; i < N; i++) {
		sum += X[i] * Y[i];
	}
	return sum;
}