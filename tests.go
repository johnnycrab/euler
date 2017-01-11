package main

import (
	"fmt"
)


func main() {
	for i := 0; i<11; i++ {
		fmt.Println((i*i*i*i -5*i*i + 5)%11)
	}
}