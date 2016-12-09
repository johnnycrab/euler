package main

import "fmt"


func main() {

	// n is the congruence class of d
	/*for n := 0; n < 840; n++ {

		// d = 3 mod 8
		if n%8 == 3 {
			if n%3 == 1 {
				if (n*n)%5 == 4 {
					if (n*n*n)%7 == 1 {
						fmt.Println(n)	
					}
					
				}
			}
		}
	}*/

	for n := 0; n<11; n++ {
		if (n*n-3)%11 == 0 {
			fmt.Println(n)
		}
	}

}