#include "benchmark.h"
#include <assert.h>

float SerialSdot(int N, float X[], int incX, float Y[], int incY) {
	float sum = 0.0;

	assert(incX == 1 && incY == 1);

	for (int i = 0; i < N; i++) {
		sum += X[i] * Y[i];
	}
	return sum;
}

float SerialSdotGeneric(int N, float X[], int incX, float Y[],
			int incY) {
	float sum = 0.0;
	int xi = 0, yi = 0;

	while(N--) {
		sum += X[xi] * Y[yi];
		xi += incX;
		yi += incY;
	}

	return sum;
}
