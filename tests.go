package main

import (
	"fmt"
)


func main() {
	var a float64 = float64(519432)/float64(632382)
	
	var b float64 = a

	for i := 1; i < 525806; i++ {
		b *= a
	}

	fmt.Println(b)
}