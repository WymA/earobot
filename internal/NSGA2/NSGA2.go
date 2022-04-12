package snga

const (
    ParetoDominating int = 0 
    ParetoDominated 	= 1
    ParetoNondominated  = 2
    ParetoEqual=3
)


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

func compare( ind1 *NSGA2Ind, ind2 *NSGA2Ind ) int {

     better := false;
     worse := false;

    for  i := 0; !( worse && better ) && ( i < kObjNum ) ; i++ {

        if ind1.y_var[i] < ind2.y_var[i] {
			better = true
		}
            
        if ind2.y_var[i] < ind1.y_var[i] {
			worse = true
		}
            
    }

    if worse {

        if better {
            return ParetoNondominated;
		}        else {
            return ParetoDominated;
		}

    }else{

        if better {
            return ParetoDominating
		}        else {
            return ParetoEqual
		}
    }
}

func fastNondominatedSort() {
	// for  i := 0 ; i < pareto_front.size() ; i++ {
	// 	pareto_front[i].clear() ;
	// }

	//pareto_front.clear();

	nextFront []NSGA2Ind
	index []int


	for  i := 1 ; !population.empty() ; i++ {

		// nextFront.clear() ;
		// index.clear();
		for p := 0 ; p < population.size() ; p++ {

			population[p].counter = 0 ;

			for q := 0 ; q < population.size()  ; q++  {

				int res = compare( population[p].indiv, population[q].indiv ) ;
				if ( kParetoDominated == res )
					population[p].counter++ ;
			}

			if ( population[p].counter == 0 ){

				population[p].rank = i ;
				nextFront.push_back( population[p] ) ;
				index.push_back(p);
			}

		}
		for ( int idx = index.size()-1 ; idx >= 0 ; idx-- )
			population.erase( population.begin() + index[idx] ) ;

		pareto_front.push_back( nextFront ) ;

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
