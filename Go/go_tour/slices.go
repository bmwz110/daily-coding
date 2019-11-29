package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sli := make([][]uint8, dy)
	for i := range sli {
		sli[i] = make([]uint8, dx)
		for j := range sli[i] {
			sli[i][j] = uint8(i*j)
		}
	
	}
	return sli
}

func main() {
	pic.Show(Pic)
}
