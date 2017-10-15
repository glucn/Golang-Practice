package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result:=make([][]uint8,dy)
	
	for y:=range result {
		result[y]=make([]uint8, dx)
		for x:=range result[y] {
			result[y][x] = uint8(x*y)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
