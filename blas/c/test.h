#ifndef __TEST_H__
#define __TEST_H__

/* Function declaration */
void SerialSaxpy(int N, float alpha, float X[], float Y[]);
void SerialSaxpyGeneric(int N, float alpha, float X[], int incX, float Y[],
						int incY);
void ISPCSaxpy(int N, float alpha, float X[], int incX, float Y[], int incY);

double SerialDasum(int N, double X[], int incX);
double SerialDasumGeneric(int N, double X[], int incX);
double ISPCDasum(int N, double X[], int incX);
typedef double (*dasumFn)(int N, double X[], int incX);

float SerialSdot(int N, float X[], int incX, float Y[], int incY);
float SerialSdotGeneric(int N, float X[], int incX, float Y[], int incY);
float ISPCSdot(int N, float X[], int incX, float Y[], int incY);
typedef float (*sdotFn)(int N, float X[], int incX, float Y[], int incY);

float SerialSdsdot(int N, float alpha, float X[], int incX, float Y[],
				int incY);
float SerialSdsdotGeneric(int N, float alpha, float X[], int incX, float Y[],
				int incY);
float ISPCSdsdot(int N, float alpha, float X[], int incX, float Y[],
				int incY);
typedef float (*sdsdotFn)(int N, float alpha, float X[], int incX, float Y[],
				int incY);

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

void SerialDasumTest(test_t *test);
void SerialDasumGenericTest(test_t *test);
void ISPCDasumTest(test_t *test);

void SerialSdotTest(test_t *test);
void SerialSdotGenericTest(test_t *test);
void ISPCSdotTest(test_t *test);

void SerialSdsdotTest(test_t *test);
void SerialSdsdotGenericTest(test_t *test);
void ISPCSdsdotTest(test_t *test);

void SerialMandelbrotTest(test_t *test);
void ISPCMandelbrotTest(test_t *test);

void SleepTest(test_t *test);

/* Test declarations */
#endif
