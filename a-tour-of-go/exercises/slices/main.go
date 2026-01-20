package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for y := range dy {
		img[y] = make([]uint8, dx)

		for x := range dx {
			img[y][x] = uint8(x ^ y)
			// img[y][x] = uint8((x + y) / 2)
			// img[y][x] = uint8(x * y)
		}
	}

	return img
}

func main() {
	pic.Show(Pic)
}
