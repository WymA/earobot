package NSGA2

import (
	"earobot/internal/common"
	"sort"
	"time"
)

var population NSGA2Population
var offspring NSGA2Population
var evolutionaryAlgo common.EvolutionaryAlgo

var evalFunc []func(ind *NSGA2Ind) float64

type NSGA2Ind struct {
	common.Individual
	Rank           int
	DominatedCount int

	// SectorialAngle float64
	// SectorialIndex int

}

type NSGA2Population struct {
	Individuals []NSGA2Ind
	ParetoFront [][]NSGA2Ind
}

func Init(initEvolutionaryAlgo common.EvolutionaryAlgo, initEvalFunc []func(ind *NSGA2Ind) float64) {

	evolutionaryAlgo = initEvolutionaryAlgo
	evalFunc = initEvalFunc

	if len(evalFunc) != evolutionaryAlgo.ObjectivesNumber {
		panic("Init NSGA-II failed, the number of evaluation functions should equal to the numbers of objectives ")
	}

	initPopulation()

}

func initPopulation() {

	population.Individuals = make([]NSGA2Ind, evolutionaryAlgo.PopulationSize)

	for i := 0; i < len(population.Individuals); i++ {

		population.Individuals[i].RandomInit(evolutionaryAlgo.VariablesLength, evolutionaryAlgo.VariablesMin, evolutionaryAlgo.VariablesMax)
		population.Individuals[i].Objectives = make([]float64, evolutionaryAlgo.ObjectivesNumber)
	}

	population.ParetoFront = make([][]NSGA2Ind, 0)
}

func geneCrossover() {

}

func geneMutation() {

}

func crowdDistAssign() {

	for _, front := range population.ParetoFront {

		if len(front) > 0 {

			frontSize := len(front)

			for _, ind := range front {
				ind.CrowdDist = 0.0
			}

			for objectiveIdx := 0; objectiveIdx < evolutionaryAlgo.ObjectivesNumber; objectiveIdx++ {

				sort.SliceStable(front, func(i, j int) bool {
					return front[i].Objectives[objectiveIdx] < front[j].Objectives[objectiveIdx]
				})

				front[0].CrowdDist, front[frontSize-1].CrowdDist = common.INFINITE, common.INFINITE

				for i := 1; i < frontSize-1; i++ {
					front[i].CrowdDist += front[i+1].Objectives[objectiveIdx] - front[i-1].Objectives[objectiveIdx]
				}

			}
		}
	}

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
func evaluation() (count int, timeConsume time.Duration, highestResult float64, lowestResult float64) {

	count = 0
	startTime := time.Now()
	for _, v := range population.Individuals {

		for i := 0; i < evolutionaryAlgo.ObjectivesNumber; i++ {
			v.Objectives[i] = evalFunc[i](&v)

			if v.Objectives[i] > highestResult {
				highestResult = v.Objectives[i]
			}
			if v.Objectives[i] < lowestResult {
				lowestResult = v.Objectives[i]
			}
			count++
		}
	}

	timeConsume = time.Since(startTime)

	return count, timeConsume, highestResult, lowestResult

}

// one generation
func runOneGerration() {

	// create offspring generation
	//geneCrossover()
	//geneMutation()

	//offspring = evolutionaryAlgo.CreateOffspring(population.Individuals.Individual)

	evaluation()

	fastNondominatedSort()

	// population clear
	population.Individuals = make([]NSGA2Ind, evolutionaryAlgo.PopulationSize)

	//for i := 0; (i < pareto_front.size()) && (population.size()+pareto_front[i].size() < pop_size); i++ {
	for i := 0; len(population.Individuals)+len(population.ParetoFront[0]) < evolutionaryAlgo.PopulationSize; i++ {
		// 	crowdDistAssign(pareto_front[i])
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

}

// Evolve begins
func Evolve() {

	for i := 0; i < evolutionaryAlgo.TotalGeneration; i++ {
		runOneGerration()
	}
}
