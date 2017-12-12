export uniform double ISPCDasum(uniform int N, uniform double X[], uniform int incX) {
	double sum = 0;

	foreach (i = 0 ... N) {
		double x = X[i];
		if( x < 0 ){
			x = -x;
		}
		sum += x;
	}
	return reduce_add(sum);
}

