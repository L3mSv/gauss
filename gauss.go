package main

import (
	"fmt"
	"math/rand"
	"time"
)

// constants
const MAXN = 2000
const max = 5
const min = 0

// global variables
var N int // matrix dimension
var A [MAXN][MAXN]float64
var B [MAXN]float64
var X [MAXN]float64

// creates the A matrix and the B vector
func initializeInputs() {
	//fmt.Println("\nInitializing...")

	for col := 0; col < N; col++ {
		for row := 0; row < N; row++ {
			A[row][col] = rand.Float64() * (max - min) // putting the number into an interval
		}
		B[col] = rand.Float64() * (max - min) // putting the number into an interval
		X[col] = 0.0
	}
}

// function to print matrices A and B
func printInputs() {
	if N < 10 {
		fmt.Println("\nMatrix A:")
		for row := 0; row < N; row++ {
			for col := 0; col < N; col++ {
				//if col < N-1 {
				//	fmt.Printf("%5.2f, ", A[row][col])
				//	continue
				//}
				//fmt.Printf("%5.2f;\n", A[row][col])

				//or this, that is infact smaller:
				fmt.Printf("%5.2f%s", A[row][col], ternary(col < N-1, ", ", ";\n"))
			}
		}

		fmt.Println("\nVector B:")
		fmt.Print("[")
		for col := 0; col < N; col++ {
			//if col < N-1 {
			//	fmt.Printf("%5.2f; ", B[col])
			//	continue
			//}
			//fmt.Printf("%5.2f]\n", B[col])

			// or this, that is infact smaller:
			fmt.Printf("%5.2f%s", B[col], ternary(col < N-1, "; ", "]\n"))
		}
	}
}

// function to print the solution vector X
func printX() {
	if N < 100 {
		fmt.Println("\nSolution Vector X:")
		fmt.Print("[")
		for row := 0; row < N; row++ {
			//if row < N-1 {
			//	fmt.Printf("%5.2f; ", X[row])
			//	continue
			//}
			//fmt.Printf("%5.2f]\n", X[row])

			//or this, that is infact smaller:
			fmt.Printf("%5.2f%s", X[row], ternary(row < N-1, "; ", "]\n"))

		}
	}
}

// function to be the "ternary" of C
func ternary(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

func gauss() {

	// gauss elimination (forward elimination)
	for norm := 0; norm < N-1; norm++ {
		for row := norm + 1; row < N; row++ {
			multiplier := A[row][norm] / A[norm][norm]
			for col := norm; col < N; col++ {
				A[row][col] -= A[norm][col] * multiplier
			}
			B[row] -= B[norm] * multiplier
		}
	}

	printInputs()

	// backforward substitution
	for row := N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col := N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}

func execute() {
	start := time.Now()
	defer func() { // runs when the function ends/return
		fmt.Println("Execution time:", time.Since(start))
	}()

	// calculate the gauss elimination
	gauss()

	return
}

func main() {

	fmt.Print("Enter matrix dimension N (1 to ", MAXN, "): ")
	fmt.Scan(&N)

	if N < 1 || N > MAXN {
		fmt.Printf("N = %d is out of range.\n", N)
		return
	}

	// inicialize A and B
	initializeInputs()

	// prints the matrix A
	printInputs()

	//calculate the time to execute the gauss method of solving systems
	execute()
	//start := time.Now()
	//gauss()
	//fmt.Println("Execution time:", time.Since(start))

	// print the solution
	printX()
}
