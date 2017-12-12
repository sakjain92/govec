#include "benchmark.h"

double SerialDasum(int N, double X[], int incX) {
	double sum = 0.0;

	if(incX != 1) {
		return sum;
	}

	for(int i = 0; i < N; i++) {
		double x = X[i];
		if(x < 0) {
			x = -x;
		}
		sum += x;
	}

	return sum;
}
