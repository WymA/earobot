package common

type Individual struct {
	Variables  []float64
	Objectives []float64
}

type Population struct {
}

func GeneCrossInd(crossRate float64, parent1 *Individual, parent2 *Individual, child1 *Individual, child2 *Individual) {

	randomRate := RandomUtil.Float64() //%1000 / 100

	if randomRate < crossRate {
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
}
