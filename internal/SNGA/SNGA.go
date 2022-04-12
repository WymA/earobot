package snga

type SNGAInd struct {
}

// //uptate the extreme point up/down
// func update_extreme_point(ind *SNGAInd) {

// }

// func init_population() {

// }
// func Pareto_HyperVolume_compare_sectorialgrid(ind *SNGAInd) {

// }
// func GetFastigiateHyperVolume(ind *SNGAInd, indIndex int, ReferencePoint []float64) {

// }

// func compute_hypervolume(mysectorpop []SNGAInd, mypopsize int, mynobj int) {

// }
// func tour_selection_hv_difference(p int, mypopulation []SNGAInd) {

// }
// func tour_selection_hv2(mypopulation []SNGAInd) {

// }

// func uniform_selection(ind_selected *SNGAInd) {

// }
// func reset_angle() {

// }

// func population2front(mypopulation []SNGAInd, population_front [][]float64) {

// }
// func calc_distance() {

// } //calculate the average distance between ps and current solutions

func genSelection() {

}

func genMutation() {

}

func evaluation() {

}

func fastNondominatedSort() {

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
