package testfunction

import (
	"earobot/internal/NSGA2"
	"earobot/internal/common"
	"testing"
)

func Test_function1(t *testing.T) {
	NSGA2.Init(common.EvolutionaryAlgo{
		VariablesLength:  3,
		VariablesMin:     -5.0,
		VariablesMax:     5.0,
		ObjectivesNumber: 2,
		CroseoverRate:    0.5,
		MutationRate:     0.5,
		PopulationSize:   100,
		TotalGeneration:  100,
	},
		[]func(ind *NSGA2.NSGA2Ind) float64{
			NSGA2Func1,
			NSGA2Func2,
		})

	NSGA2.Evaluation()

}
