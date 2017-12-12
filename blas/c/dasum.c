#include "benchmark.h"
#include <assert.h>

double SerialDasum(int N, double X[], int incX) {
	double sum = 0.0;
  int i;
	assert(incX == 1);

	for(i = 0; i < N; i++) {
		double x = X[i];
		if(x < 0) {
			x = -x;
		}
		sum += x;
	}

	return sum;
}

double SerialDasumGeneric(int N, float alpha, double X[], int incX) {
	double sum = 0.0;
	int xi = 0;

	while(N--) {
		double x = X[xi];
		if(x < 0) {
			x = -x;
		}
		sum += x;
		xi += incX;
	}

	return sum;
}
