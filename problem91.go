/*
	Problem 91
*/

package main

import "fmt"

type Point struct {
	x,y int
}

/*

	p1
	| \
  v1|  \ v3
	|   \
	|    \
	0------p2
       v2
	Check if scalar product of the given vectors is 0
*/
func isRightAngleTriangle(p1, p2 Point) bool {
	// v1, v2
	if p1.x*p2.x + p1.y*p2.y == 0 {
		return true
	}

	// v1, v3
	if p1.x*(p2.x - p1.x) + p1.y*(p2.y - p1.y) == 0 {
		return true
	}

	// v2, v3
	if p2.x*(p2.x - p1.x) + p2.y*(p2.y - p1.y) == 0 {
		return true
	}

	return false
}

func main() {
	count := 0

	// just iterate over all points and form triangles without duplicates
	for y1 := 0; y1 <= 50; y1++ {
		for x1 := 0; x1 <= 50; x1++ {
			for y2 := y1; y2 <= 50; y2++ {
				for x2 := 0; x2 <= 50; x2++ {
					if y2 == y1 && x2 <= x1 {
						continue
					}

					p1 := Point{x1, y1}
					p2 := Point{x2, y2}

					if ((p1.x != 0) || (p1.y != 0)) && ((p2.x != 0) || (p2.y != 0)) && ((p1.x != p2.x) || (p1.y != p2.y)) {
						if isRightAngleTriangle(p1, p2) {
							count++
						}
					}
				}
			}			
		}
	}

	fmt.Println(count)
}