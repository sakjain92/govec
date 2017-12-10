#ifndef __TEST_H__
#define __TEST_H__

/* Function declaration */
void SerialSaxpy(int N, float alpha, float X[], float Y[]);
void SerialSaxpyGeneric(int N, float alpha, float X[], int incX, float Y[],
						int incY);
void ISPCSaxpy(int N, float alpha, float X[], float Y[]);

void SerialSaxpyTest(test_t *test);
void SerialSaxpyGenericTest(test_t *test);
void ISPCSaxpyTest(test_t *test);

/* Test declarations */
#endif
