package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy*dx)
	for i := 0; i < dy; i++ {
		img[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			img[i][j] = uint8(i^j) //Interesting functions include (x+y)/2, x*y, and x^y. 
		}

	}
	fmt.Printf("%v\n", img)
	return img
}

func main() {
	pic.Show(Pic)
}

