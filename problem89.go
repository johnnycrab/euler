/*
	Problem 89
*/

package main

import (
	"fmt"
	"bufio"
    "log"
    "os"
)

func romanStyle(n int, modulo int, current string, symbol string, nineSymbol string, fiveSymbol string, fourSymbol string) (int, string) {
	numOfSymbols := n/modulo

	if fiveSymbol != "" {
		if numOfSymbols >= 5 {
			numOfSymbols -= 5
			n -= 5*modulo
			current += fiveSymbol
		}

		if numOfSymbols == 4 {
			numOfSymbols -= 4
			n -= 4*modulo
			current += fourSymbol
		}
	}

	for i:=0; i<numOfSymbols; i++ {
		current += symbol
	}

	n -= numOfSymbols * modulo

	if nineSymbol != "" && n >= (modulo/10)*9 {
		current += nineSymbol
		n -= 9 * (modulo/10)
	}

	return n, current
}

func integerToRoman(n int) string {
	roman := ""

	n, roman = romanStyle(n, 1000, roman, "M", "CM", "", "")
	n, roman = romanStyle(n, 100, roman, "C", "XC", "D", "CD")
	n, roman = romanStyle(n, 10, roman, "X", "IX", "L", "XL")
	n, roman = romanStyle(n, 1, roman, "I", "", "V", "IV")

	return roman
}

func romanToInteger(roman string) int {
	sum := 0

	// split it up
	symbols := make([]string, len(roman))
	for i, v := range roman {
		symbols[i] = fmt.Sprintf("%c", v)
	}

	i := -1
	for i < len(symbols) - 1 {
		i++

		s := symbols[i]
		if s == "M" {
			sum += 1000
		} else if s == "D" {
			sum += 500
		} else if s == "C" {
			if i+1 < len(symbols) && symbols[i+1] == "M" {
				sum += 900
				i++
			} else if i+1 < len(symbols) && symbols[i+1] == "D" {
				sum += 400
				i++
			} else {
				sum += 100
			}
		} else if s == "L" {
			sum += 50
		} else if s == "X" {
			if i+1 < len(symbols) && symbols[i+1] == "C" {
				sum += 90
				i++
			} else if i+1 < len(symbols) && symbols[i+1] == "L" {
				sum += 40
				i++
			} else {
				sum += 10
			}
		} else if s == "V" {
			sum += 5
		} else if s == "I" {
			if i+1 < len(symbols) && symbols[i+1] == "X" {
				sum += 9
				i++
			} else if i+1 < len(symbols) && symbols[i+1] == "V" {
				sum += 4
				i++
			} else {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	full := 0
	reduced := 0


	file, err := os.Open("romans.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        roman := scanner.Text()
        full += len(roman)
        reduced += len(integerToRoman(romanToInteger(roman)))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(full - reduced)
}