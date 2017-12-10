#include "benchmark.h"

float SerialSdsdot(int N, float alpha, float X[], float Y[]) {
	double sum = 0.0;
	for (int i = 0; i < N; i++) {
		sum += (double)X[i] * (double)Y[i];
	}
	return (float)((double)alpha + sum);
}

float SerialSdsdotGeneric(int N, float alpha, float X[], int incX, float Y[],
						int incY) {
	double sum = 0.0;
	int xi =0, yi =0;

	while(N--) {
		sum += (double)X[xi] * (double)Y[yi];
		xi += incX;
		yi += incY;
	}

	return (float)((double)alpha + sum);
}
