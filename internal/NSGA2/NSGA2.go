package snga

import (
	"earobot/internal/common"
)

const (
	ParetoDominating   int = 0
	ParetoDominated        = 1
	ParetoNondominated     = 2
	ParetoEqual            = 3
)

var population NSGA2Population
var ObjectiveNumber int = 2

type NSGA2Ind struct {
	common.Individual
	Rank           int
	CrowdDist      float64
	DominatedCount int

	// SectorialAngle float64
	// SectorialIndex int

}

func (rcvr *NSGA2Ind) RandomInit(varRange int) {

}

// func (rcvr *NSGA2Ind) ObjEvaluation() {

// }

type NSGA2Population struct {
	Individuals []NSGA2Ind
	ParetoFront [][]NSGA2Ind
}

// //uptate the extreme point , that's f up/down
// func update_extreme_point(ind *NSGA2Ind) {

// }

// func Pareto_HyperVolume_compare_sectorialgrid(ind *NSGA2Ind) {

// }
// func GetFastigiateHyperVolume(ind *NSGA2Ind, indIndex int, ReferencePoint []float64) {

// }

// func uniform_selection(ind_selected *NSGA2Ind) {

// }

func genSelection() {

}

func genMutation() {

}

// compare two individuals
func tournament(ind1 *NSGA2Ind, ind2 *NSGA2Ind) *NSGA2Ind {

	res := compare(ind1, ind2)

	if res == ParetoDominating {
		return ind1
	}
	if res == ParetoDominated {
		return ind2
	}
	if ind1.CrowdDist >= ind2.CrowdDist {
		return ind1
	} else {
		return ind2
	}
}

// func population2front(mypopulation []NSGA2Ind, population_front [][]float64) {

// }

func compare(ind1 *NSGA2Ind, ind2 *NSGA2Ind) int {

	better := false
	worse := false

	for i := 0; !(worse && better) && (i < ObjectiveNumber); i++ {

		if ind1.Objectives[i] < ind2.Objectives[i] {
			better = true
		}

		if ind2.Objectives[i] < ind1.Objectives[i] {
			worse = true
		}

	}

	if worse {

		if better {
			return ParetoNondominated
		} else {
			return ParetoDominated
		}

	} else {

		if better {
			return ParetoDominating
		} else {
			return ParetoEqual
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

				res := compare(&population.Individuals[p], &population.Individuals[q])
				if ParetoDominated == res {
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

func evaluation() {

}

// one generation
func runOneGerration() {

	genSelection()
	genMutation()

	// population.insert(population.end(), offspring.begin(), offspring.end())

	evaluation()

	fastNondominatedSort()

	// population.clear()

	// for i := 0; (i < pareto_front.size()) && (population.size()+pareto_front[i].size() < pop_size); i++ {

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

	// }

	// GetBestObj(kObjNodes)
	// GetBestObj(kObjEnergy)

}

func execute(run int, hvl []float64, igd []float64, totaltime *float64, unevolvetime *float64) {

}
