package main

import (
	"fmt"
	"math/cmplx"
	"strings"
)

func calculateMandelbrot(
	maxIters int,
	xMin float64,
	xMax float64,
	yMin float64,
	yMax float64,
	width int,
	height int,
) [][]int {

	rows := make([][]int, height)
	for imgY := 0; imgY < height; imgY++ {
		row := make([]int, width)
		for imgX := 0; imgX < width; imgX++ {
			xPercent := float64(imgX) / float64(width)
			yPercent := float64(imgY) / float64(height)
			cx := xMin + (xMax-xMin)*xPercent
			cy := yMin + (yMax-yMin)*yPercent
			escapedAt := mandelbrotAtPoint(cx, cy, maxIters)
			row[imgX] = escapedAt
		}
		rows[imgY] = row
	}
	return rows
}

func mandelbrotAtPoint(cx, cy float64, maxIters int) int {
	z := complex(0.0, 0.0)
	c := complex(cx, cy)

	for i := 0; i < maxIters; i++ {
		if cmplx.Abs(z) > 2.0 {
			return i
		}
		z = z*z + c
	}
	return maxIters
}

func renderMandelbrot(escapeVals [][]int) {
	fmt.Println("")
	for _, row := range escapeVals {
		line := []string{}
		for _, column := range row {
			switch {
			case column >= 0 && column <= 2:
				line = append(line, " ")
			case column >= 3 && column <= 5:
				line = append(line, ".")
			case column >= 6 && column <= 10:
				line = append(line, "•")
			case column >= 11 && column <= 30:
				line = append(line, "*")
			case column >= 31 && column <= 100:
				line = append(line, "+")
			case column >= 101 && column <= 200:
				line = append(line, "x")
			case column >= 201 && column <= 400:
				line = append(line, "$")
			case column >= 401 && column <= 700:
				line = append(line, "#")
			default:
				line = append(line, "%")
			}
		}

		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println("")
}

func main() {
	mandelbrot := calculateMandelbrot(1000, -2.0, 1.0, -1.0, 1.0, 100, 24)
	renderMandelbrot(mandelbrot)
}
