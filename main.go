// Package infinity provides functions for working with infinite numbers.
package infinity

import (
	"math"
)

var (
	Positive = math.Inf(1.0)  // +∞
	Negative = math.Inf(-1.0) // -∞
)
