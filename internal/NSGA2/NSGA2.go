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

func init_population() {

}
func Pareto_HyperVolume_compare_sectorialgrid(ind *NSGA2Ind) {

}
func GetFastigiateHyperVolume(ind *NSGA2Ind, indIndex int, ReferencePoint []float64) {

}

func compute_hypervolume(mysectorpop []NSGA2Ind, mypopsize int, mynobj int) {

}
func tour_selection_hv_difference(p int, mypopulation []NSGA2Ind) {

}
func tour_selection_hv2(mypopulation []NSGA2Ind) {

}

func uniform_selection(ind_selected *NSGA2Ind) {

}
func reset_angle() {

}

func population2front(mypopulation []NSGA2Ind, population_front [][]float64) {

}
func calc_distance() {

} //calculate the average distance between ps and current solutions

// one generation
func runOneGerration() {

	// GenSelection()
	// GenMutation()

	// population.insert(population.end(), offspring.begin(), offspring.end())

	// Evaluation()

	// FastNondominatedSort()

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

	// // QString output ;

	// // output.append( "Gen "+ QString::number(cur_gen) +": " );
	// // output.append( " Nodes=" + QString::number( best_ind.y_var[kObjNodes] ) ) ;
	// // output.append( " Energy=" + QString::number( best_ind.y_var[kObjEnergy] ) ) ;
	// // output.append( " Converage: " + QString::number(maxc) ) ;

	// cur_gen++

	// // return output ;

}

func execute(run int, hvl []float64, igd []float64, totaltime *float64, unevolvetime *float64) {

}
