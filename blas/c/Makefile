ISPC=ispc
CC=gcc
CFLAGS=-I. -Wall -Werror -O3 -lrt -lm
IFLAGS=--target=sse4-x2 --arch=x86-64 -O3
#IFLAGS=--target=avx --arch=x86-64 -O3
DEPS = benchmark.h test.h Makefile

CTESTS=saxpy.o dasum.o sdot.o sdsdot.o test.o mandelbrot.o
ITESTS=saxpy_ispc.o dasum_ispc.o sdot_ispc.o sdsdot_ispc.o mandelbrot_ispc.o

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

%.o: %.ispc $(DEPS)
	$(ISPC) -o $@ $< $(IFLAGS)

benchmark: benchmark.o $(ITESTS) $(CTESTS)
	gcc -o $@ $^ $(CFLAGS)

clean:
	rm *.o
