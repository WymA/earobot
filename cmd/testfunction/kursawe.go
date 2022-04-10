package testfunction

import "math"

// x => [-5,5], i=> [1,3]
func function1(x []float64) (s float64) {

	for i := 0; i < len(x)-1; i++ {
		s += -10 * math.Exp(-0.2*math.Sqrt(math.Pow(x[i], 2.0)+math.Pow(x[i+1], 2.0)))
	}
	return s
}

// x => [-5,5], i=> [1,3]
func function2(x []float64) (s float64) {

	for i := 0; i < len(x)-1; i++ {
		s += math.Pow(math.Abs(x[i]), .8) + 5*math.Sin(math.Pow(x[i], 3.0))
	}
	return s
}
