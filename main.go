package main

import "github.com/SamKnightGit/web-pso/pso"

func main() {
	pso.PSO(100, 100, []float64{4.5, 4.5}, []float64{-4.5, -4.5}, 0.5, 0.1, 0.2, 0.8, true)
}
