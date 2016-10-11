/*
	Problem:
	https://projecteuler.net/problem=66

	This problem means we have to find fundamental solutions to Pell equations.
	We will do this by using the sequence of convergents of the regular continued fraction of sqrt(D)
*/

package main

import (
	"fmt"
	"math"
	"math/big"
)

func findMinimalSolution(_D int64) *big.Int {

	const prec = 200

	D 	  := big.NewInt(_D)
	DFloat := new(big.Float).SetInt(D)

	// find the convergents of the regular continued fraction of sqrt(D)
	h_n_2 := big.NewInt(0)
	h_n_1 := big.NewInt(1)
	x 	  := big.NewInt(0)
	k_n_2 := big.NewInt(1)
	k_n_1 := big.NewInt(0)
	y 	  := big.NewInt(0)
	cfInt := big.NewInt(0)

	tempInt := new(big.Int)
	tempFloat := new(big.Float).SetPrec(200)
	// continued fraction expansion
	sqrt := new(big.Float).SetPrec(200).SetFloat64(math.Sqrt(float64(_D)))
	sqrt.Int(cfInt)
	cfFloat := new(big.Float).SetInt(cfInt)

	m := new(big.Float).SetPrec(200).SetFloat64(0.0)
	d := new(big.Float).SetPrec(200).SetFloat64(1.0)

	for {
		// calculate convergents

		tempInt.Mul(cfInt, h_n_1)
		
		x.Add(tempInt, h_n_2)
		tempInt.Mul(cfInt, k_n_1)
		y.Add(tempInt, k_n_2)

		// move forward

		h_n_2.Set(h_n_1)
		h_n_1.Set(x)
		k_n_2.Set(k_n_1)
		k_n_1.Set(y)

		// check if we have a solution
		Res := big.NewInt(0)
		Res2 := big.NewInt(0)
		Res.Mul(x,x)
		Res2.Mul(y,y)
		Res2.Mul(Res2, D)
		Res.Sub(Res, Res2)

		if Res.Int64() == 1 {
			
			return x
		}

		// new continued fractions
		tempFloat.Mul(cfFloat, d)
		m.Sub(tempFloat, m)

		tempFloat.Mul(m, m)
		tempFloat.Sub(DFloat, tempFloat)
		d.Quo(tempFloat, d)

		tempFloat.Add(sqrt, m)
		tempFloat.Quo(tempFloat, d)

		tempFloat.Int(cfInt)
		cfFloat.SetInt(cfInt)

	}

	return nil
}



func main() {

	var solution int64 = 0 // value of D, for which the largest fundamental solution is obtained
	solutionX := big.NewInt(0)
	temp := big.NewInt(0)

	var D int64 = 1

	for D <= 1000 {
		// if D is square, we don't have any positive integer solutions
		isSquare := false
		
		var i int64 = 1
		for i*i<=D {
			if i*i == D {
				isSquare = true
				break
			}
			i++
		}

		if isSquare {
			D++
			continue
		}

		minSol := findMinimalSolution(D)
		//fmt.Println("solution for D: ", D, " x: ", minSol)
		
		if temp.Sub(solutionX, minSol).Sign() == -1 {
			solutionX.Set(minSol)
			solution = D
		}

		D++
	}

	fmt.Println(solution)
}