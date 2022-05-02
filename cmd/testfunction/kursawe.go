package testfunction

import (
	"earobot/internal/NSGA2"
	"math"
)

// x => [-5,5], i=> [1,3]
func NSGA2Func1(ind *NSGA2.NSGA2Ind) (s float64) {

	for i := 0; i < len(ind.Variables)-1; i++ {
		s += -10 * math.Exp(-0.2*math.Sqrt(math.Pow(ind.Variables[i], 2.0)+math.Pow(ind.Variables[i+1], 2.0)))
	}
	return s
}

// x => [-5,5], i=> [1,3]
func NSGA2Func2(ind *NSGA2.NSGA2Ind) (s float64) {

	for i := 0; i < len(ind.Variables)-1; i++ {
		s += math.Pow(math.Abs(ind.Variables[i]), .8) + 5*math.Sin(math.Pow(ind.Variables[i], 3.0))
	}
	return s
}
