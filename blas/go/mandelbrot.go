//+build !ASSEMBLY

package blas

import (
    "github.com/sakjain92/govectool/govec"
)

func mandel(c_re float32 , c_im float32 , count int) int {
    var z_re float32
    var z_im float32

    z_re = c_re
    z_im = c_im

    var i int
    for i = 0; i < count; i++ {

        var new_re, new_im float32

        if (z_re * z_re + z_im * z_im > 4.0) {
            break
        }

        new_re = z_re*z_re - z_im*z_im
        new_im = 2.0 * z_re * z_im

        z_re = c_re + new_re
        z_im = c_im + new_im
    }
    return i
}

func SerialMandelbrot(
    x0 float32, y0 float32, x1 float32, y1 float32,
    width int, height int,
    startRow int, totalRows int,
    maxIterations int,
    output []int) {

    var dx, dy float32
    var endRow, i, j int

    dx = (x1 - x0) / float32(width)
    dy = (y1 - y0) / float32(height)

    endRow = startRow + totalRows

    for j = startRow; j < endRow; j++ {

        for i = 0; i < width; i++ {
            var x, y float32
            var index int

            x = x0 + float32(i) * dx
            y = y0 + float32(j) * dy

            index = (j * width + i)
            output[index] = mandel(x, y, maxIterations)
        }
    }
}

func __govecMandel(c_re float32 , c_im float32 , count int) int {
    var z_re float32
    var z_im float32

    z_re = c_re
    z_im = c_im

    var i int
    for i = 0; i < count; i++ {

        var new_re, new_im float32

        if z_re * z_re + z_im * z_im > 4.0 {
            break
        }

        new_re = z_re*z_re - z_im*z_im
        new_im = 2.0 * z_re * z_im

        z_re = c_re + new_re
        z_im = c_im + new_im
    }
    return i
}

func _govecMandelbrot(
            x0 govec.UniformFloat32, y0 govec.UniformFloat32,
            x1 govec.UniformFloat32, y1 govec.UniformFloat32,
            width govec.UniformInt, height govec.UniformInt,
            startRow govec.UniformInt, totalRows govec.UniformInt,
            maxIterations govec.UniformInt,
            output []govec.UniformInt) {

    var dx, dy float32
    var i, j int

    dx = (float32)(x1 - x0) / (float32)(width)
    dy = (float32)(y1 - y0) / (float32)(height)

    for govec.DoubleRange(j, 0, height, i, 0, width) {
        var x, y float32
        var index int

        x = (float32)(x0) + (float32)(i) * dx
        y = (float32)(y0) + (float32)(j) * dy

        index = (j * (int)(width) + i)
        output[index] = (govec.UniformInt)(__govecMandel(x, y, (int)(maxIterations)))
    }
}

