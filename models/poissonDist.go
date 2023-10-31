package models

import (
	"math"
	"math/rand"
)

type PoissonDist struct {
}

func NewPoissonDist() *PoissonDist {
	return &PoissonDist{}
}

func (pd *PoissonDist) Generate(lambda float64) int {
	u := rand.Float64()

	p := 0.0
	for i := 0; i <= int(lambda); i++ {
		p += math.Pow(lambda/math.Exp(lambda), float64(i)) / float64(factorial(i))
		if u < p {
			return i
		}
	}

	return int(lambda)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
