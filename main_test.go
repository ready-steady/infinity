package infinity

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestLinear(t *testing.T) {
	m, n := uint(4), uint(3)

	A := []float64{
		+0.0, +1.0, -2.0, +0.0,
		-1.0, -2.0, +0.0, +1.0,
		+1.0, +1.0, +0.0, +2.0,
	}

	test := func(x, y []float64) {
		assert.Equal(Linear(A, x, m, n), y, t)
	}

	test([]float64{1.0, 2.0, 1.0}, []float64{-1.0, -2.0, -2.0, 4.0})
	test([]float64{Plus, 2.0, 1.0}, []float64{-1.0, Plus, Minus, 4.0})
	test([]float64{1.0, Minus, 1.0}, []float64{Plus, Plus, -2.0, Minus})
	test([]float64{1.0, 2.0, Plus}, []float64{Plus, Plus, -2.0, Plus})
	test([]float64{Plus, 2.0, Minus}, []float64{Minus, -4.0, Minus, Minus})
}

func TestQuadratic(t *testing.T) {
	m := uint(3)

	A := []float64{
		+0.0, +1.0, -2.0,
		-1.0, -2.0, +0.0,
		+1.0, +1.0, +0.0,
	}

	test := func(x []float64, y float64) {
		assert.Equal(Quadratic(A, x, m), y, t)
	}

	test([]float64{1.0, 2.0, 3.0}, -5.0)
	test([]float64{Plus, 2.0, 3.0}, Minus)
	test([]float64{1.0, Plus, 3.0}, Minus)
	test([]float64{1.0, 2.0, Plus}, Plus)
	test([]float64{1.0, 1.0, Plus}, -2.0)
	test([]float64{1.0, 1.0, Minus}, -2.0)
}
