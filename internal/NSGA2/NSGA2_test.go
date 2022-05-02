package NSGA2

import (
	"earobot/internal/common"
	"math"
	"testing"
)

// x => [-5,5], i=> [1,3]
func NSGA2Func1(ind *NSGA2Ind) (s float64) {

	for i := 0; i < len(ind.Variables)-1; i++ {
		s += -10 * math.Exp(-0.2*math.Sqrt(math.Pow(ind.Variables[i], 2.0)+math.Pow(ind.Variables[i+1], 2.0)))
	}
	return s
}

// x => [-5,5], i=> [1,3]
func NSGA2Func2(ind *NSGA2Ind) (s float64) {

	for i := 0; i < len(ind.Variables)-1; i++ {
		s += math.Pow(math.Abs(ind.Variables[i]), .8) + 5*math.Sin(math.Pow(ind.Variables[i], 3.0))
	}
	return s
}

func Test_function(t *testing.T) {

	Init(common.EvolutionaryAlgo{
		VariablesLength:  3,
		VariablesMin:     -5.0,
		VariablesMax:     5.0,
		ObjectivesNumber: 2,
		CroseoverRate:    0.5,
		MutationRate:     0.5,
		PopulationSize:   100,
		TotalGeneration:  100,
	},
		[]func(ind *NSGA2Ind) float64{
			NSGA2Func1,
			NSGA2Func2,
		})
	count, timeConsume, highestResult, lowestResult := evaluation()
	t.Logf("\ncount: %d, timeConsume: %d, highestResult: %f, lowestResult: %f", count, timeConsume, highestResult, lowestResult)
}

func TestInit(t *testing.T) {
	type args struct {
		initEvolutionaryAlgo common.EvolutionaryAlgo
		initEvalFunc         []func(ind *NSGA2Ind) float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.initEvolutionaryAlgo, tt.args.initEvalFunc)
		})
	}
}

func Test_crowdDistAssign(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crowdDistAssign()
		})
	}
}
