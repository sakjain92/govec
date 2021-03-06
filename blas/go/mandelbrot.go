//+build !ASSEMBLY

package blas

//go:generate govectool mandelbrot.go

import (
    "github.com/sakjain92/govectool/govec"
    "github.com/harrydb/go/img/pnm"
    "os"
    "log"
    "image"
    "image/color"
)

func mandel(c_re float32 , c_im float32 , count int) int32 {
    var z_re float32
    var z_im float32

    z_re = c_re
    z_im = c_im

    var i int32
    for i = 0; i < (int32)(count); i++ {

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
    output []int32) {

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

func __govecISPCMandel(c_re float32 , c_im float32 , count int) int32 {
    var z_re float32
    var z_im float32

    z_re = c_re
    z_im = c_im

    var i int32
    for i = 0; i < (int32)(count); i++ {

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

func _govecISPCMandelbrot(
            x0 govec.UniformFloat32, y0 govec.UniformFloat32,
            x1 govec.UniformFloat32, y1 govec.UniformFloat32,
            width govec.UniformInt, height govec.UniformInt,
            startRow govec.UniformInt, totalRows govec.UniformInt,
            maxIterations govec.UniformInt,
            output []govec.UniformInt32) {

    var dx, dy float32
    var i, j int
    var endRow govec.UniformInt

    dx = (float32)(x1 - x0) / (float32)(width)
    dy = (float32)(y1 - y0) / (float32)(height)

    endRow = startRow + totalRows

    for govec.DoubleRange(j, startRow, endRow, i, 0, width) {
        var x, y float32
        var index int

        x = (float32)(x0) + (float32)(i) * dx
        y = (float32)(y0) + (float32)(j) * dy

        index = (j * (int)(width) + i)
        output[index] = (govec.UniformInt32)(__govecISPCMandel(x, y, (int)(maxIterations)))
    }
}

type mandelImage struct {
	output []int32
	width	int
	height	int
}

func (m mandelImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (m mandelImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m mandelImage) At(x, y int) color.Color {

	var c color.RGBA
	v := m.output[x + y * m.width]

	c.R = uint8(v)
	c.G = uint8(v)
	c.B = uint8(v)
	c.A = 1

	return c
}

// Writes a mandelbrot image to .ppm file
func writeMandelBrotFile(output []int32, width int, height int, filename string) {

	var m mandelImage

	m.output = output
	m.width = width
	m.height = height

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = pnm.Encode(file, m, pnm.PPM)
	if err != nil {
		log.Fatal(err)
	}

}
