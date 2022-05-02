package common

import "math"

const (
	ParetoDominating   int = 0
	ParetoDominated        = 1
	ParetoNondominated     = 2
	ParetoEqual            = 3
)

const (
	INFINITE float64 = math.MaxFloat64
)

type Individual struct {
	Variables  []float64
	Objectives []float64
	CrowdDist  float64
}

type EvolutionaryAlgo struct {
	VariablesLength  int
	VariablesMin     float64
	VariablesMax     float64
	ObjectivesNumber int
	CroseoverRate    float64
	MutationRate     float64
	PopulationSize   int
	TotalGeneration  int
}

func (rcvr *Individual) RandomInit(length int, min float64, max float64) {
	rcvr.Variables = make([]float64, length)
	for i := 0; i < length; i++ {
		rcvr.Variables[i] = min + RandomUtil.Float64()*max
	}
}

func (rcvr *EvolutionaryAlgo) geneCrossover(parent1 *Individual, parent2 *Individual) (child1 *Individual, child2 *Individual) {

	child1.RandomInit(rcvr.VariablesLength, rcvr.VariablesMin, rcvr.VariablesMax)
	child2.RandomInit(rcvr.VariablesLength, rcvr.VariablesMin, rcvr.VariablesMax)

	randomRate := RandomUtil.Float64() //%1000 / 100

	if randomRate < rcvr.CroseoverRate {
		// crosover gene between wo parents to
		parent1CrossPoint := 1 + RandomUtil.Intn(len(parent1.Variables)-1)
		parent2CrossPoint := 1 + RandomUtil.Intn(len(parent2.Variables)-1)

		child1.Variables = make([]float64, parent1CrossPoint)
		child2.Variables = make([]float64, parent2CrossPoint)

		copy(child1.Variables, parent1.Variables[:parent1CrossPoint])
		copy(child2.Variables, parent2.Variables[:parent2CrossPoint])

		child1.Variables = append(child1.Variables, parent2.Variables[parent2CrossPoint:]...)
		child2.Variables = append(child2.Variables, parent1.Variables[parent1CrossPoint:]...)

	} else {

		child1 = parent1
		child2 = parent2
	}

	return child1, child2
}

func (rcvr *EvolutionaryAlgo) geneMutation(ind *Individual) {

	mutationPoint := RandomUtil.Intn(len(ind.Variables))

	if RandomUtil.Float64() < rcvr.MutationRate {

		// #FIXME magic number
		if RandomUtil.Float64() < 0.5 {
			ind.Variables[mutationPoint] = RandomUtil.Float64() * (rcvr.VariablesMax - ind.Variables[mutationPoint])
		} else {
			ind.Variables[mutationPoint] = RandomUtil.Float64() * (ind.Variables[mutationPoint] - rcvr.VariablesMin)
		}
	}

}

// compare two individuals
func (rcvr *EvolutionaryAlgo) Compare(ind1 *Individual, ind2 *Individual) int {

	better := false
	worse := false

	for i := 0; !(worse && better) && (i < rcvr.ObjectivesNumber); i++ {

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

func (rcvr *EvolutionaryAlgo) tournament(ind1 *Individual, ind2 *Individual) *Individual {

	res := rcvr.Compare(ind1, ind2)

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

func (rcvr *EvolutionaryAlgo) CreateOffspring(parents []Individual) (offspring []Individual) {

	offspring = make([]Individual, rcvr.PopulationSize)

	randomParentsIndex := make([]int, rcvr.PopulationSize)

	for i := range randomParentsIndex {
		randomParentsIndex[i] = i
	}

	RandomUtil.Shuffle(len(randomParentsIndex), func(i, j int) {
		randomParentsIndex[i], randomParentsIndex[j] = randomParentsIndex[j], randomParentsIndex[i]
	})

	for i := 0; i < rcvr.PopulationSize; i += 2 {
		parent1 := rcvr.tournament(&parents[randomParentsIndex[i]], &parents[randomParentsIndex[i+1]])
		parent2 := rcvr.tournament(&parents[randomParentsIndex[i+2]], &parents[randomParentsIndex[i+3]])
		child1, child2 := rcvr.geneCrossover(parent1, parent2)
		offspring = append(offspring, *child1, *child2)
	}

	return offspring
}
