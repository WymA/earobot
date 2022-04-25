package snga

import (
	"earobot/internal/common"
)

var population NSGA2Population
var evolutionaryAlgo common.EvolutionaryAlgo
var evalFunc1 func(ind *NSGA2Ind)
var evalFunc2 func(ind *NSGA2Ind)

type NSGA2Ind struct {
	common.Individual
	Rank           int
	DominatedCount int

	// SectorialAngle float64
	// SectorialIndex int

}

// func (rcvr *NSGA2Ind) ObjEvaluation() {

// }

type NSGA2Population struct {
	Individuals []NSGA2Ind
	ParetoFront [][]NSGA2Ind
}

func initPopulation() {

	population.Individuals = make([]NSGA2Ind, evolutionaryAlgo.PopulationSize)
	population.ParetoFront = make([][]NSGA2Ind, 0)
}

func genMutation() {

}

func fastNondominatedSort() {

	population.ParetoFront = make([][]NSGA2Ind, 0)

	var nextFront []NSGA2Ind
	var index []int

	for i := 1; len(population.Individuals) > 0; i++ {

		nextFront = make([]NSGA2Ind, 0)
		index = make([]int, 0)
		for p := 0; p < len(population.Individuals); p++ {

			population.Individuals[p].DominatedCount = 0

			for q := 0; q < len(population.Individuals); q++ {

				res := evolutionaryAlgo.Compare(&population.Individuals[p].Individual, &population.Individuals[q].Individual)
				if common.ParetoDominated == res {
					population.Individuals[p].DominatedCount++
				}

			}

			if population.Individuals[p].DominatedCount == 0 {

				population.Individuals[p].Rank = i
				nextFront = append(nextFront, population.Individuals[p])
				index = append(index, p)
			}

		}
		for idx := len(index) - 1; idx >= 0; idx-- {

			population.Individuals = append(population.Individuals[:index[idx]], population.Individuals[index[idx]+1:]...) // delete element
		}

		population.ParetoFront = append(population.ParetoFront, nextFront)

	}
}

//
func Evaluation(func1 func(ind *NSGA2Ind), func2 func(ind *NSGA2Ind)) (ind *NSGA2Ind) {

	return ind
}

// one generation
func runOneGerration() {

	// GeneCrossover()
	genMutation()

	// population.insert(population.end(), offspring.begin(), offspring.end())

	Evaluation(evalFunc1, evalFunc2)

	fastNondominatedSort()

	// population.clear()

	//for i := 0; (i < pareto_front.size()) && (population.size()+pareto_front[i].size() < pop_size); i++ {
	for i := 0; len(population.Individuals)+len(population.ParetoFront[0]) < evolutionaryAlgo.PopulationSize; i++ {
		// 	CrowdDistAssign(pareto_front[i])
		// 	population.insert(population.end(), pareto_front[i].begin(), pareto_front[i].end())

		// }

		// if population.size() < pop_size {

		// 	SortByPareto(pareto_front[i])
		// 	population.insert(population.end(), pareto_front[i].begin(),
		// 		pareto_front[i].begin()+pop_size-population.size()) //Offset
		// }

		// // maxc = population.begin()->indiv.coverage ;

		// for i := 0; i < population.size(); i++ {
		// 	if population[i].indiv.coverage > maxc {
		// 		maxc = population[i].indiv.coverage
		// 	}

	}

	// GetBestObj(kObjNodes)
	// GetBestObj(kObjEnergy)

}

// Evolve begins
func Evolve() {

	initPopulation()
	for i := 0; i < evolutionaryAlgo.TotalGeneration; i++ {
		runOneGerration()
	}
}
