package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants
const MAXN = 2000

// Global variables
var N int
var A [MAXN][MAXN]float64
var B [MAXN]float64
var X [MAXN]float64

// Function to initialize random values in A and B
func initializeInputs() {
	fmt.Println("\nInitializing...")

	for col := 0; col < N; col++ {
		for row := 0; row < N; row++ {
			A[row][col] = rand.Float64() / 32768.0
		}
		B[col] = rand.Float64() / 32768.0
		X[col] = 0.0
	}
}

// Function to print matrices A and B
func printInputs() {
	if N < 10 {
		fmt.Println("\nMatrix A:")
		for row := 0; row < N; row++ {
			for col := 0; col < N; col++ {
				fmt.Printf("%5.2f%s", A[row][col], ifElse(col < N-1, ", ", ";\n"))
			}
		}

		fmt.Println("\nVector B:")
		fmt.Print("[")
		for col := 0; col < N; col++ {
			fmt.Printf("%5.2f%s", B[col], ifElse(col < N-1, "; ", "]\n"))
		}
	}
}

// Function to print the solution vector X
func printX() {
	if N < 100 {
		fmt.Println("\nSolution Vector X:")
		fmt.Print("[")
		for row := 0; row < N; row++ {
			fmt.Printf("%5.2f%s", X[row], ifElse(row < N-1, "; ", "]\n"))
		}
	}
}

// Helper function for conditional expressions
func ifElse(condition bool, a, b string) string {
	if condition {
		return a
	}
	return b
}

// Gaussian elimination without pivoting
func gauss() {
	fmt.Println("Computing Serially.")

	// Gaussian elimination (forward elimination)
	for norm := 0; norm < N-1; norm++ {
		for row := norm + 1; row < N; row++ {
			multiplier := A[row][norm] / A[norm][norm]
			for col := norm; col < N; col++ {
				A[row][col] -= A[norm][col] * multiplier
			}
			B[row] -= B[norm] * multiplier
		}
	}

	// Back substitution
	for row := N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col := N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}

func main() {
	// Input parameters
	fmt.Print("Enter matrix dimension N (1 to ", MAXN, "): ")
	fmt.Scan(&N)

	if N < 1 || N > MAXN {
		fmt.Printf("N = %d is out of range.\n", N)
		return
	}

	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Initialize A and B
	initializeInputs()

	// Print input matrices if N is small
	printInputs()

	// Start timing
	startTime := time.Now()

	// Perform Gaussian elimination
	gauss()

	// Stop timing
	elapsedTime := time.Since(startTime).Milliseconds()
	fmt.Printf("\nElapsed time = %d ms.\n", elapsedTime)

	// Print the solution vector X
	printX()
}
