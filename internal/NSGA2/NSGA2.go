package snga

type NSGA2Ind struct {
	Variables      []float64
	Objectives     []float64
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

//uptate the extreme point , that's f up/down
func update_extreme_point(ind *NSGA2Ind) {

}

func Pareto_HyperVolume_compare_sectorialgrid(ind *NSGA2Ind) {

}
func GetFastigiateHyperVolume(ind *NSGA2Ind, indIndex int, ReferencePoint []float64) {

}

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

func uniform_selection(ind_selected *NSGA2Ind) {

}

func genSelection() {

}

func genMutation() {

}

func population2front(mypopulation []NSGA2Ind, population_front [][]float64) {

}

func fastNondominatedSort() {

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
