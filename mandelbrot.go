package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"bufio"
	"os"
)

const WRITE_LOCATION string  = "C:\\Users\\mike2_000\\Desktop\\derp.png"
const SQUARE_RADIUS  float64 = 2.0
const IMAGE_SIZE     int     = 800
const MAX_ITERATIONS int     = 10000

func num_iterations_to_escape(c complex128, max_iterations int) int {
	z := complex(0, 0)
	for iteration := 0; iteration < max_iterations; iteration++ {
		z *= z
		z += c
		if cmplx.Abs(z) >= 2.0 {
			return iteration
		}
	}
	return max_iterations
}

func get_mandelbrot(square_radius float64, size int, max_iterations int) image.Image {
	hSize := size / 2;
	img := image.NewRGBA(image.Rect(-hSize, -hSize, hSize, hSize))
	for i := -hSize; i <= hSize; i++ {
		for j := -hSize; j <= hSize; j++ {
			scale              := square_radius / float64(hSize)
			c                  := complex(scale * float64(i), scale * float64(j))
			num_iterations     := num_iterations_to_escape(c, max_iterations)
			iterations_percent := float64(max_iterations - num_iterations) / float64(max_iterations)
			greyscale          := uint8(float64(255) * iterations_percent)
			color              := color.RGBA{greyscale, greyscale, greyscale, 255}
			img.Set(i, j, color)
		}
	}
	return img
}
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_image(img image.Image, location string) {
	f, err := os.Create(location)
	check(err)
	writer := bufio.NewWriter(f)
	err = png.Encode(writer, img)
	check(err)
}
		
func main() {
	img := get_mandelbrot(SQUARE_RADIUS, IMAGE_SIZE, MAX_ITERATIONS)
	write_image(img, WRITE_LOCATION)
}
