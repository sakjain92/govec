
void SerialSaxpy(int N, float alpha, float X[], float Y[]) {
	for (int i = 0; i < N; i++) {
		Y[i] += alpha * X[i];
	}
}

void SerialSaxpyGeneric(int N, float alpha, float X[], int incX, float Y[],
						int incY) {

	if (incX == 1 && incY == 1) {
		for (int i = 0; i < N; i++) {
			Y[i] += alpha * X[i];
		}
		return;
	}

	int xi, yi;

	switch alpha {
	case 0: 
		break;
	case 1:
		for (; N >= 2; N -= 2) {
			Y[yi] += X[xi];
			xi += incX;
			yi += incY;

			Y[yi] += X[xi];
			xi += incX;
			yi += incY;
		}
		if (N != 0) {
			Y[yi] += alpha * X[xi];
		}
		break;
	case -1:
		for (; N >= 2; N -= 2) {
			Y[yi] -= X[xi];
			xi += incX;
			yi += incY;

			Y[yi] -= X[xi];
			xi += incX;
			yi += incY;
		}
		if (N != 0) {
			Y[yi] -= X[xi];
		}
		break;
	default:
		for (; N >= 2; N -= 2) {
			Y[yi] += alpha * X[xi];
			xi += incX;
			yi += incY;

			Y[yi] += alpha * X[xi];
			xi += incX;
			yi += incY;
		}
		if (N != 0) {
			Y[yi] += alpha * X[xi];
		}
	}
}
