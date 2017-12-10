export void ISPCSaxpy(uniform int N,
                      uniform float alpha,
	                  uniform float X[],
	                  uniform float Y[])
{
    foreach (i = 0 ... N) {           
        Y[i] = alpha * X[i] + Y[i];
    }
}
