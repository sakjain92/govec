export void govecSaxpy(uniform int N, uniform float alpha, uniform float X[],
		 uniform int incX, uniform float Y[], uniform int incY) {

	foreach ( i = 0 ... N){
//		Y[i * incY] += alpha * X[i * incX];
		Y[i] += alpha * X[i];
	};
}
