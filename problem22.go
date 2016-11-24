package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func bubblesort(list []string) {

	doAgain := true
	for doAgain {
		doAgain = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				doAgain = true
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}

}

func main() {
	
	bytes, err := ioutil.ReadFile("p022_names.txt")

	if err == nil {

		names := strings.Split(string(bytes), ",")

		bubblesort(names)
		
		// get name scores
		result := 0

		for i, name := range names {
			letterSum := 0
			for _, r := range name {
				if string(r) != "\"" {
					letterSum += int(r - 64)
				} 
			}

			result += (i+1) * letterSum
		}

		fmt.Println(result)
	}
}