void  govecSaxpy(int N, float alpha, float * X, int incX, float * Y, int incY) {

	int xi, yi, i;

	xi = 0;
	yi = 0;

	for( i = 0; i < N; i++ ){
		Y[yi] += alpha * X[xi];
		xi += incX;
		yi += incY;
	};
}
