// Package infinity provides functions for working with infinite numbers.
package infinity

import (
	"math"

	"github.com/ready-steady/linear/matrix"
)

var (
	Plus  = math.Inf(1.0)  // +∞
	Minus = math.Inf(-1.0) // -∞
)

// Linear multiplies an m-by-n matrix by an n-element vector. The elements of
// the vector can potentially be infinite.
func Linear(A, x []float64, m, n uint) []float64 {
	y := make([]float64, m)
	ok, s := inspect(x, n)
	if ok {
		matrix.Multiply(A, x, y, m, n, 1)
		return y
	}
	for i := uint(0); i < m; i++ {
		fin, inf := 0.0, 0.0
		for j := uint(0); j < n; j++ {
			a := A[j*m+i]
			if a == 0.0 {
				continue
			}
			if s[j] == 0.0 {
				fin += a * x[j]
			} else {
				inf += a * s[j]
			}
		}
		if inf != 0.0 {
			y[i] = inf * Plus
		} else {
			y[i] = fin
		}
	}
	return y
}

// Quadratic multiplies an m-by-m matrix by an m-element vector from both sides.
// The elements of the vector can potentially be infinite.
func Quadratic(A, x []float64, m uint) float64 {
	ok, s := inspect(x, m)
	if ok {
		y := make([]float64, m)
		matrix.Multiply(A, x, y, m, m, 1)
		return matrix.Dot(x, y, m)
	}
	Fin, Inf, INF := 0.0, 0.0, 0.0
	for i := uint(0); i < m; i++ {
		fin, inf := 0.0, 0.0
		for j := uint(0); j < m; j++ {
			a := A[j*m+i]
			if a == 0.0 {
				continue
			}
			if s[j] == 0.0 {
				fin += a * x[j]
			} else {
				inf += a * s[j]
			}
		}
		if s[i] == 0.0 {
			Fin += x[i] * fin
			Inf += x[i] * inf
		} else {
			Inf += s[i] * fin
			INF += s[i] * inf
		}
	}
	if INF != 0.0 {
		return INF * Plus
	} else if Inf != 0.0 {
		return Inf * Plus
	} else {
		return Fin
	}
}

func inspect(x []float64, m uint) (bool, []float64) {
	ok, signs := true, make([]float64, m)
	for i := uint(0); i < m; i++ {
		switch x[i] {
		case Plus:
			ok, signs[i] = false, +1.0
		case Minus:
			ok, signs[i] = false, -1.0
		}
	}
	return ok, signs
}
