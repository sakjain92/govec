export void ISPCSdsdot(uniform int N,
                       uniform float alpha,
	                   uniform float X[],
	                   uniform float Y[])
{
    double sum = 0.0;
    foreach (i = 0 ... N) {           
        sum += X[i] * Y[i];
    }
    return (float)((double)alpha + reduce_add(sum));
}
