#include <stdio.h>
#include <math.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include "benchmark.h"
#include "test.h"
#include "time.h"

/* Number of samples per test */
#define MAX_SAMPLES	1000

/* Name of test case */
#define MAX_NAME	100

#define USE_CLOCK_GETTIME

typedef struct stats {
	double min;
	double max;
	double median;
	double avg;
	double stddev;
} stat_t;

typedef struct out {
	double time[MAX_SAMPLES];
	stat_t stats;
} out_t;

typedef struct test {
	char name[MAX_NAME];
	void (*func)(struct test *test);
	out_t outval;
	unsigned long long startTicks;
	unsigned long long stopTicks;
	int stopped;
	int itr;		/* Iteration number */
} test_t;

#define DEF_TEST(x) {#x, x ##Test, {{0}, {0}}, 0, 0, 0}
test_t tests[] = {
	DEF_TEST(SerialSaxpy),
	DEF_TEST(SerialSaxpyGeneric),
	DEF_TEST(ISPCSaxpy),
	DEF_TEST(SerialMandelbrot),
	DEF_TEST(ISPCMandelbrot),
	//DEF_TEST(Sleep),
};

int cmpsort(const void *p1, const void *p2)
{
	return *((double *)p1) - *((double *)p2);
}

void computeStats(out_t *outval)
{
	int i;
	double min = 1E20, max = 0;
	double sum = 0, avg, stdsum = 0;

	// sort first for median
	qsort(&outval->time, MAX_SAMPLES, sizeof(double), cmpsort);

	for (i = 0; i < MAX_SAMPLES; i++) {

		double val = outval->time[i];

		if (val > max) max = val;
		if (val < min) min = val;
		sum += val;
	}

	outval->stats.min = min;
	outval->stats.max = max;
	outval->stats.avg = avg = sum / (double)MAX_SAMPLES;

	if (MAX_SAMPLES % 2) {
		outval->stats.median = outval->time[MAX_SAMPLES / 2];
	} else {
		outval->stats.median = (outval->time[MAX_SAMPLES / 2] +
					outval->time[MAX_SAMPLES / 2 - 1]) / 2;
	}

	for (i = 0; i < MAX_SAMPLES; i++) {

		double val = outval->time[i];
		stdsum += fabs(val - avg) * fabs(val -avg);
	}

	outval->stats.stddev = sqrt(stdsum / (double)MAX_SAMPLES);
}

void printStats(out_t *outval)
{
	stat_t *stats = &outval->stats;
	char *unit;
	double d;

	if (stats->avg > 1E6) {
		unit = "ms";
		d = 1E6;
	} else if (stats->avg > 1E3) {
		unit = "us";
		d = 1E3;
	} else {
		unit = "ns";
		d = 1;
	}

	fprintf(stdout, "Min %.2f %s,\t Max %.2f %s,\t Avg: %.2lf %s,\t Median: %.2f %s\t Stddev: %.2lf %s\n",
			stats->min / d, unit, stats->max / d, unit, stats->avg / d, unit, stats->median / d, unit,
			stats->stddev / d, unit);
}

unsigned long long getTicks(void)
{
#ifndef USE_CLOCK_GETTIME
	unsigned int a, d;
	asm volatile("rdtsc" : "=a" (a), "=d" (d));
	return (unsigned long long)(a) | ((unsigned long long)(d) << 32);
#else
	struct timespec tp;
	assert(clock_gettime(CLOCK_MONOTONIC, &tp) == 0);

	return tp.tv_sec * 1E9 + tp.tv_nsec;
#endif
}

double secondsPerTick(void)
{
	FILE *fp = fopen("/proc/cpuinfo","r");
	double secondsPerTick_val;

	char input[1024];

	if (!fp) {
		fprintf(stderr, "failed: couldn't find /proc/cpuinfo.");
		exit(-1);

	}

	secondsPerTick_val = 1e-9;
	while (!feof(fp) && fgets(input, 1024, fp)) {

		float GHz, MHz;
		if (strstr(input, "model name")) {

			char* at_sign = strstr(input, "@");
			if (at_sign) {
				char* after_at = at_sign + 1;
				char* GHz_str = strstr(after_at, "GHz");
				char* MHz_str = strstr(after_at, "MHz");
				if (GHz_str) {
					*GHz_str = '\0';
					if (1 == sscanf(after_at, "%f", &GHz)) {
						secondsPerTick_val = 1e-9f / GHz;
						goto found;
					}
				} else if (MHz_str) {
					*MHz_str = '\0';
					if (1 == sscanf(after_at, "%f", &MHz)) {
						secondsPerTick_val = 1e-6f / GHz;
						goto found;
					}
				}
			}

		} else if (1 == sscanf(input, "cpu MHz : %f", &MHz)) {
			secondsPerTick_val = 1e-6f / MHz;
			goto found;
		}
	}

	fprintf(stderr, "failed: couldn't find Processor speed");
	exit(-1);

found:
	fclose(fp);
	return secondsPerTick_val;
}

void resetTime(test_t *test)
{
	test->startTicks = getTicks();
}

void stopTime(test_t *test)
{
	if (test->stopped == 0) {
		test->stopTicks = getTicks();
		test->stopped = 1;
	}
}

int isFirst(test_t *test)
{
	return test->itr == 0;
}

static double getTestTimens(test_t *test, double secPerTick)
{
	unsigned long long ticks = test->stopTicks - test->startTicks;

	assert(test->stopTicks > test->startTicks);

#ifndef USE_CLOCK_GETTIME
	return (double)(ticks) * secPerTick * 1E9;
#else
	return (double)(ticks);
#endif
}

static void initTest(test_t *test, int itr)
{
	resetTime(test);
	test->stopped = 0;
	test->itr = itr;
}

static void finishTest(test_t *test)
{
	stopTime(test);
}

int main(void)
{
	int i, j;
	int num_test = sizeof(tests) / sizeof(tests[0]);
	test_t *test;

	double secPerTick = secondsPerTick();

	fprintf(stdout, "Process Speed: %f Ghz\n", 1E-9/secPerTick);
	fprintf(stdout, "No. of samples taken: %d\n", MAX_SAMPLES);

	for (i = 0; i < num_test; i++) {

		test = &tests[i];

		fprintf(stdout, "######## Running test %s [%d/%d] ########\n",
				test->name, i + 1, num_test);

		memset(test->outval.time, 0, MAX_SAMPLES * sizeof(long));

		for (j = 0; j < MAX_SAMPLES; j++) {

			initTest(test, j);
			tests[i].func(test);
			finishTest(test);
			test->outval.time[j] = getTestTimens(test, secPerTick);
		}

		computeStats(&test->outval);
		printStats(&test->outval);
	}

	fprintf(stdout, "****** All test completed *******\n");

	return 0;
}
