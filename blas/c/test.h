#ifndef __TEST_H__
#define __TEST_H__

/* Function declaration */
void SerialSaxpy(int N, float alpha, float X[], float Y[]);
void SerialSaxpyGeneric(int N, float alpha, float X[], int incX, float Y[],
						int incY);
void ISPCSaxpy(int N, float alpha, float X[], int incX, float Y[], int incY);

void SerialMandelbrot(
    float x0, float y0, float x1, float y1,
    int width, int height,
    int startRow, int totalRows,
    int maxIterations,
    int output[]);

void ISPCMandelbrot( float x0,  float y0,
                             float x1,  float y1,
                             int width,  int height,
			     int startRow, int totalRows,
                             int maxIterations,
                             int output[]);

typedef void (*mandelbrotFn)(float x0,  float y0,
                             float x1,  float y1,
                             int width,  int height,
			     int startRow, int totalRows,
                             int maxIterations,
                             int output[]);

void SerialSaxpyTest(test_t *test);
void SerialSaxpyGenericTest(test_t *test);
void ISPCSaxpyTest(test_t *test);

void SerialMandelbrotTest(test_t *test);
void ISPCMandelbrotTest(test_t *test);

void SleepTest(test_t *test);

/* Test declarations */
#endif
