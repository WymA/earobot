package snga

type SNGAInd struct {
}

//uptate the extreme point , that's f up/down
func update_extreme_point(ind *SNGAInd) {

}

func init_population() {

}
func Pareto_HyperVolume_compare_sectorialgrid(ind *SNGAInd) {

}
func GetFastigiateHyperVolume(ind *SNGAInd, indIndex int, ReferencePoint []float64) {

}

func compute_hypervolume(mysectorpop []SNGAInd, mypopsize int, mynobj int) {

}
func tour_selection_hv_difference(p int, mypopulation []SNGAInd) {

}
func tour_selection_hv2(mypopulation []SNGAInd) {

}
func execute(run int, hvl []float64, igd []float64, totaltime *float64, unevolvetime *float64) {

}
func uniform_selection(ind_selected *SNGAInd) {

}
func reset_angle() {

}
func save_population(mypopulation []SNGAInd, saveFilename string) {

}
func save_ps(saveFilename string) {

}

func population2front(mypopulation []SNGAInd, population_front [][]float64) {

}
func calc_distance() {

} //calculate the average distance between ps and current solutions

// Execute the algorithm
func Execute() {

}
