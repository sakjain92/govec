#include "benchmark.h"
#include "test.h"
#include <stdlib.h>
#include <assert.h>
#include <unistd.h>

enum saxpyFunc {
	SERIAL = 0,
	SERIAL_GENERIC = 1,
	SERIAL_ISPC = 2
};

float frand()
{
	return (float)(rand())/ (float)(RAND_MAX);
}

/* Saxpy */
static void SaxpyTest(enum saxpyFunc num, test_t *test)
{
	int i, count = 100000;
	float X[count], Y[count], res[count];
	float alpha;
	int incX = 1, incY = 1;

	(void)i;
	(void)res;
	alpha = 1;

	int first = isFirst(test);
	if (first) {
		alpha = frand();

		for (i = 0; i < count; i++) {
			X[i] = Y[i] = frand();
			res[i] = X[i] * alpha + Y[i];
		}
	}

	switch (num) {
	case SERIAL:
		resetTime(test);
		SerialSaxpy(count, alpha, X, Y);
		stopTime(test);
		break;
	case SERIAL_GENERIC:
		resetTime(test);
		SerialSaxpyGeneric(count, alpha, X, incX, Y, incY);
		stopTime(test);
		break;
	case SERIAL_ISPC:
		resetTime(test);
		ISPCSaxpy(count, alpha, X, incX, Y, incY);
		stopTime(test);
		break;
	default:
		assert(0);
	}

	if (first) {
		for (i = 0; i < count; i++) {
			assert(Y[i] == res[i]);
		}
	}
}

void SerialSaxpyTest(test_t *test)
{
	SaxpyTest(SERIAL, test);
}

void SerialSaxpyGenericTest(test_t *test)
{
	SaxpyTest(SERIAL_GENERIC, test);
}

void ISPCSaxpyTest(test_t *test)
{
	SaxpyTest(SERIAL_ISPC, test);
}

void SleepTest(test_t *test)
{
	sleep(1);
}
