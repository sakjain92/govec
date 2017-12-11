#ifndef __BENCHMARK_H__
#define __BENCHMARK_H__

struct test;
typedef struct test test_t;

void resetTime(test_t *test);
void stopTime(test_t *test);
int isFirst(test_t *test);

#endif
