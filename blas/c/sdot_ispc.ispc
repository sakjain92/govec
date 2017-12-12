export uniform float ISPCSdot(uniform int N, uniform float X[], uniform int incX, uniform float Y[], uniform int incY) {
	float sum = 0.0;
	foreach ( i = 0 ... N ) {
		sum += X[i] * Y[i];
	}
	return reduce_add(sum);
}

